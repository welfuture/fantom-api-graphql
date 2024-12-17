package rpc

import (
	"encoding/json"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
	"math/big"
	"net/http"
	"net/mail"
	"net/url"
	"regexp"
	"time"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/sti.json --pkg contracts --type StakerInfo --out ./contracts/staker_info.go

// stiRequestTimeout is number of seconds we wait for the staker information request to finish.
const stiRequestTimeout = 1 * time.Second

// stiNameCheckRegex is the expression used to check for staker name validity
var stiNameCheckRegex = regexp.MustCompile(`^[\w\s.\-_'$()]+$`)

// StakerInfo extracts an extended staker information from smart contact by their id.
func (ftm *FtmBridge) StakerInfo(id *hexutil.Big) (*types.StakerInfo, error) {
	if id == nil {
		return nil, fmt.Errorf("staker ID not given")
	}

	// keep track of the operation
	ftm.log.Debugf("loading staker information for staker #%d", id.ToInt().Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewStakerInfo(ftm.sfcConfig.StiContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate STI contract: %v", err)
		return nil, err
	}

	// call for data
	stUrl, err := contract.GetInfo(nil, (*big.Int)(id))
	if err != nil {
		ftm.log.Errorf("failed to get the staker #%d information: %v", id.ToInt().Uint64(), err)
		return nil, err
	}

	// var url string
	if len(stUrl) == 0 {
		ftm.log.Debugf("no information for staker #%d", id.ToInt().Uint64())
		return nil, nil
	}

	// try to download JSON for the info
	return ftm.downloadStakerInfo(stUrl)
}

// downloadStakerInfo tries to download staker information from the given URL address.
func (ftm *FtmBridge) downloadStakerInfo(stUrl string) (*types.StakerInfo, error) {
	// log what we are about to do
	ftm.log.Debugf("downloading staker info address [%s]", stUrl)

	// make a http client
	cl := http.Client{Timeout: stiRequestTimeout}

	// prep request
	req, err := http.NewRequest(http.MethodGet, stUrl, nil)
	if err != nil {
		ftm.log.Errorf("can not request given staker info at [%s]; %s", stUrl, err.Error())
		return nil, err
	}

	// be honest, set agent
	req.Header.Set("User-Agent", "Fantom GraphQL API Server")

	// process the request
	res, err := cl.Do(req)
	if err != nil {
		ftm.log.Errorf("can not download staker info from [%s]; %s", stUrl, err.Error())
		return nil, err
	}

	// read the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		ftm.log.Errorf("can not read staker info response from [%s]; %s", stUrl, err.Error())
		return nil, err
	}

	// try to parse
	var info types.StakerInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		ftm.log.Errorf("can not decode staker info [%s]; %s", stUrl, err.Error())
		ftm.log.Debugf("Received: %s", body)
		return nil, err
	}

	// do we have anything?
	if !ftm.isValidStakerInfo(&info) {
		ftm.log.Errorf("invalid staker info content [%s]", stUrl)
		return nil, err
	}

	ftm.log.Debugf("found staker [%s]", *info.Name)
	return &info, nil
}

// isValidStakerInfo check if the staker information is valid and can be used.
func (ftm *FtmBridge) isValidStakerInfo(info *types.StakerInfo) bool {
	// name must be available
	if nil == info.Name || 0 == len(*info.Name) || !stiNameCheckRegex.Match([]byte(*info.Name)) {
		ftm.log.Error("staker name not valid")
		return false
	}

	// check the logo URL
	if !isValidStakerInfoUrl(info.LogoUrl, true) {
		ftm.log.Error("staker logo URL not valid")
		return false
	}

	// check the website
	if !isValidStakerInfoUrl(info.Website, false) {
		ftm.log.Error("staker website URL not valid")
		return false
	}

	// check the contact URL
	if !isValidStakerInfoUrl(info.Contact, false) && !isValidEmailContact(info.Contact) {
		ftm.log.Error("staker contact URL not valid link or email address")
		return false
	}
	return true
}

// isValidStakerInfoUrl validates the given URL address from the staker info.
func isValidStakerInfoUrl(addr *string, reqHttps bool) bool {
	// do we even have a URL; it's ok if not
	if nil == addr || 0 == len(*addr) {
		return true
	}

	// try to decode the address
	u, err := url.ParseRequestURI(*addr)
	if err != nil || u.Scheme == "" || (reqHttps && u.Scheme != "https") {
		return false
	}
	return true
}

// isValidEmailContact validates contact as an email address.
func isValidEmailContact(addr *string) bool {
	_, err := mail.ParseAddress(*addr)
	return err == nil
}
