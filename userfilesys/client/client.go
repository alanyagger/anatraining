package client

// CS 161 Project 2
// ANA-Training Task 4

// You MUST NOT change these default imports. ANY additional imports
// may break the autograder!

import (
	"encoding/json"

	userlib "github.com/cs161-staff/project2-userlib"
	"github.com/google/uuid"

	// "encoding/hex" is permitted to be an optional import
	// hex.EncodeToString(...) is useful for converting []byte to string

	// Useful for string manipulation
	"strings"

	// Useful for formatting strings (e.g. `fmt.Sprintf`).
	//"fmt"

	// Useful for creating new error messages to return using errors.New("...")
	"errors"

	// Optional.
	_ "strconv"
)

// This serves two purposes: it shows you a few useful primitives,
// and suppresses warnings for imports not being used. It can be
// safely deleted!
/*func someUsefulThings() {

	// Creates a random UUID.
	randomUUID := uuid.New()

	// Prints the UUID as a string. %v prints the value in a default format.
	// See https://pkg.go.dev/fmt#hdr-Printing for all Golang format string flags.
	userlib.DebugMsg("Random UUID: %v", randomUUID.String())

	// Creates a UUID deterministically, from a sequence of bytes.
	hash := userlib.Hash([]byte("user-structs/alice"))
	deterministicUUID, err := uuid.FromBytes(hash[:16])
	if err != nil {
		// Normally, we would `return err` here. But, since this function doesn't return anything,
		// we can just panic to terminate execution. ALWAYS, ALWAYS, ALWAYS check for errors! Your
		// code should have hundreds of "if err != nil { return err }" statements by the end of this
		// project. You probably want to avoid using panic statements in your own code.
		panic(errors.New("An error occurred while generating a UUID: " + err.Error()))
	}
	userlib.DebugMsg("Deterministic UUID: %v", deterministicUUID.String())

	// Declares a Course struct type, creates an instance of it, and marshals it into JSON.
	type Course struct {
		Name      string
		professor []byte
	}

	course := Course{"CS 161", []byte("Nicholas Weaver")}
	courseBytes, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}

	userlib.DebugMsg("Struct: %v", course)
	userlib.DebugMsg("JSON Data: %v", courseBytes)

	// Generate a random private/public keypair.
	// The "_" indicates that we don't check for the error case here.
	var pk userlib.PKEEncKey
	var sk userlib.PKEDecKey
	pk, sk, _ = userlib.PKEKeyGen()
	userlib.DebugMsg("PKE Key Pair: (%v, %v)", pk, sk)

	// Here's an example of how to use HBKDF to generate a new key from an input key.
	// Tip: generate a new key everywhere you possibly can! It's easier to generate new keys on the fly
	// instead of trying to think about all of the ways a key reuse attack could be performed. It's also easier to
	// store one key and derive multiple keys from that one key, rather than
	originalKey := userlib.RandomBytes(16)
	derivedKey, err := userlib.HashKDF(originalKey, []byte("mac-key"))
	if err != nil {
		panic(err)
	}
	userlib.DebugMsg("Original Key: %v", originalKey)
	userlib.DebugMsg("Derived Key: %v", derivedKey)

	// A couple of tips on converting between string and []byte:
	// To convert from string to []byte, use []byte("some-string-here")
	// To convert from []byte to string for debugging, use fmt.Sprintf("hello world: %s", some_byte_arr).
	// To convert from []byte to string for use in a hashmap, use hex.EncodeToString(some_byte_arr).
	// When frequently converting between []byte and string, just marshal and unmarshal the data.
	//
	// Read more: https://go.dev/blog/strings

	// Here's an example of string interpolation!
	_ = fmt.Sprintf("%s_%d", "file", 1)
}*/

// This is the type definition for the User struct.
// A Go struct is like a Python or Java class - it can have attributes
// (e.g. like the Username attribute) and methods (e.g. like the StoreFile method below).
type User struct {
	Username          string
	Password          string
	UUID              uuid.UUID
	Invitor_storeykey uuid.UUID
	Shared_filename   string
	// You can add other attributes here if you want! But note that in order for attributes to
	// be included when this struct is serialized to/from JSON, they must be capitalized.
	// On the flipside, if you have an attribute that you want to be able to access from
	// this struct's methods, but you DON'T want that value to be included in the serialized value
	// of this struct that's stored in datastore, then you can use a "private" variable (e.g. one that
	// begins with a lowercase letter).
}

// NOTE: The following methods have toy (insecure!) implementations.

func InitUser(username string, password string) (userdataptr *User, err error) {
	var userdata User
	userdata.Username = username
	userdata.Password = password
	userdata.Shared_filename = ""
	if username == "" {
		return nil, errors.New(strings.ToTitle("username is empty"))
	}
	deterministicUUID, err := uuid.FromBytes(userlib.Hash([]byte(username))[:16])
	if err != nil {
		return nil, err
	}
	userdata.UUID = deterministicUUID
	_, ok := userlib.DatastoreGet(userdata.UUID)
	if ok {
		return nil, errors.New(strings.ToTitle("user has existed"))
	}
	userbytes, err := json.Marshal(&userdata)
	if err != nil {
		return nil, err
	}
	userlib.DatastoreSet(userdata.UUID, userbytes)
	return &userdata, nil
}

func GetUser(username string, password string) (userdataptr *User, err error) {
	var userdata User
	var tempdata []byte
	hash := userlib.Hash([]byte(username))
	testUUID, _ := uuid.FromBytes(hash[:16])
	tempdata, ok := userlib.DatastoreGet(testUUID)
	if !ok {
		return nil, errors.New(strings.ToTitle("user does not exist"))
	}
	err = json.Unmarshal(tempdata, &userdata)
	if err != nil {
		return nil, err
	}
	if userdata.Password != password {
		return nil, errors.New(strings.ToTitle("password is not correct"))
	}
	userdataptr = &userdata
	return userdataptr, nil
}

func (userdata *User) StoreFile(filename string, content []byte) (err error) {
	storageKey, err := uuid.FromBytes(userlib.Hash([]byte(filename + userdata.Username))[:16])
	if err != nil {
		return err
	}
	contentBytes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	userlib.DatastoreSet(storageKey, contentBytes)
	return nil
}

func (userdata *User) AppendToFile(filename string, content []byte) error {
	var precontentBytes, precontent []byte
	storageKey, err := uuid.FromBytes(userlib.Hash([]byte(filename + userdata.Username))[:16])
	if err != nil {
		return err
	}
	if filename == userdata.Shared_filename {
		storageKey = userdata.Invitor_storeykey
	}
	precontentBytes, ok := userlib.DatastoreGet(storageKey)
	if !ok {
		return errors.New(strings.ToTitle("file not found"))
	}
	err = json.Unmarshal(precontentBytes, &precontent)
	if err != nil {
		return err
	}
	mergedcontent := append(precontent, content...)
	mergedcontentBytes, err := json.Marshal(mergedcontent)
	if err != nil {
		return err
	}
	userlib.DatastoreSet(storageKey, mergedcontentBytes)
	return nil
}

func (userdata *User) LoadFile(filename string) (content []byte, err error) {
	storageKey, err := uuid.FromBytes(userlib.Hash([]byte(filename + userdata.Username))[:16])
	if err != nil {
		return nil, err
	}
	if filename == userdata.Shared_filename {
		storageKey = userdata.Invitor_storeykey
	}
	dataJSON, ok := userlib.DatastoreGet(storageKey)
	if !ok {
		return nil, errors.New(strings.ToTitle("file not found"))
	}
	err = json.Unmarshal(dataJSON, &content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (userdata *User) CreateInvitation(filename string, recipientUsername string) (invitationPtr uuid.UUID, err error) {
	u := uuid.New()
	invitationPtr, err = uuid.FromBytes(userlib.Hash([]byte(filename + recipientUsername))[:16])
	if err != nil {
		return u, err
	}
	owner_storageKey, err := uuid.FromBytes(userlib.Hash([]byte(filename + userdata.Username))[:16])
	if err != nil {
		return u, err
	}
	if filename == userdata.Shared_filename {
		owner_storageKey = userdata.Invitor_storeykey
	}
	owner_storageKey_Bytes, err := json.Marshal(owner_storageKey)
	if err != nil {
		return u, err
	}
	userlib.DatastoreSet(invitationPtr, owner_storageKey_Bytes)
	return
}

func (userdata *User) AcceptInvitation(senderUsername string, invitationPtr uuid.UUID, filename string) error {
	userdata.Shared_filename = filename
	var temp uuid.UUID
	rawstoreykey, ok := userlib.DatastoreGet(invitationPtr)
	if !ok {
		return errors.New(strings.ToTitle("invitation not found"))
	}
	err := json.Unmarshal(rawstoreykey, &temp)
	if err != nil {
		return err
	}
	userdata.Invitor_storeykey = temp
	return nil
}

func (userdata *User) RevokeAccess(filename string, recipientUsername string) error {
	revokedPtr, err := uuid.FromBytes(userlib.Hash([]byte(filename + recipientUsername))[:16])
	if err != nil {
		return err
	}
	userlib.DatastoreDelete(revokedPtr)

	return nil
}
