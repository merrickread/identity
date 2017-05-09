package main

import (
  "errors"
  "fmt"
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "encoding/json"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Identity struct {
  IDHash               string   `json:"idHash"`
  DeviceFingerPrint    string   `json:"deviceFingerPrint"`
  FirstName            string   `json:"firstName"`
  LastName             string   `json:"lastName"`
  GovernmentID         string   `json:"governmentId"`
  PhoneNumber          string   `json:"phoneNumber"`
  Email                string   `json:"email"`
  TransactionType      string   `json:"transactionType"`
  Amount               string   `json:"amount"`
  Time                 string   `json:"timestamp"`
  BankNumber           string   `json:"bankNumber"`
}

type DeviceTransaction struct {
  DeviceFingerPrint    string  `json:"deviceFingerPrint"`
  IDHash               string  `json:"idHash"`
  TransactionType      string  `json:"transactionType"`
  Amount               string  `json:"amount"`
  Time                 string  `json:"timestamp"`
  Location             string  `json:"location"`
  Status               bool    `json:"status"`
}

func main() {
  err := shim.Start(new(SimpleChaincode))
  if err != nil {
    fmt.Printf("Error starting Simple chaincode: %s", err)
  }
}

// Init to deploy the chaincode with our identity struct
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

  // var identity Identity
  // bytes, err := json.Marshal(identity)
  // if err != nil { return nil, errors.New("Error creating Identity record") }

  // pass single user ID hash as arg
  err := stub.PutState("identity", []byte(args[0]))

  if err != nil {
    return nil, err
  }

  return nil, nil
}

func (t *SimpleChaincode) create_identity(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

  var i Identity

   // idHash              := "\"IDHash\":\""+args[0]+"\", "
   // deviceFingerPrint   := "\"DeviceFingerPrint\":\""+args[1]+"\", "
   // firstName           := "\"FirstName\":\""+args[2]+"\", "
   // lastName            := "\"LastName\":\""+args[3]+"\", "
   // governmentId        := "\"GovernmentId\":\""+args[4]+"\", "
   // phoneNumber         := "\"PhoneNumber\":\""+args[5]+"\", "
   // email               := "\"Email\":\""+args[6]+"\", "
   // bankNumber          := "\"BankNumber\":\""+args[7]+"\""

  // variables to define json object for go
  // this is weird but seems to be only way I can figure it out for now

  idHash              := "\"IDHash\":\""+args[0]+"\", "
  deviceFingerPrint   := "\"DeviceFingerPrint\":\""+args[1]+"\", "
  firstName           := "\"FirstName\":\""+args[2]+"\", "
  lastName            := "\"LastName\":\""+args[3]+"\", "
  governmentId        := "\"GovernmentId\":\""+args[4]+"\", "
  phoneNumber         := "\"PhoneNumber\":\""+args[5]+"\", "
  email               := "\"Email\":\""+args[6]+"\", "
  transactionType     := "\"TransactionType\":\""+args[7]+"\", "
  amount              := "\"Amount\":\""+args[8]+"\", "
  timestamp           := "\"Timestamp\":\""+args[9]+"\", "
  bankNumber          := "\"BankNumber\":\""+args[10]+"\""

  identity_json := "{"+idHash+deviceFingerPrint+firstName+lastName+governmentId+phoneNumber+email+transactionType+amount+timestamp+bankNumber+"}"

  // Convert the JSON defined above into a vehicle object for go
  err := json.Unmarshal([]byte(identity_json), &i)

  if err != nil { return nil, errors.New("Invalid JSON object") }

  save, err := t.save_identity(stub, i)

  if save != true { 
    fmt.Printf("CREATE_IDENTITY: Error converting identity record: %s", save);
  }

  return nil, nil
}

func (t *SimpleChaincode) save_identity(stub shim.ChaincodeStubInterface, i Identity) (bool, error) {

  bytes, err := json.Marshal(i)

  if err != nil { fmt.Printf("SAVE_CHANGES: Error converting identity record: %s", err); return false, errors.New("Error converting identity record") }

  err = stub.PutState(i.IDHash, bytes)

  if err != nil { fmt.Printf("SAVE_CHANGES: Error storing identity record: %s", err); return false, errors.New("Error storing identity record") }

  return true, nil
}

func (t *SimpleChaincode) record_transaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

  var d DeviceTransaction

  deviceFingerPrint   := "\"DeviceFingerPrint\":\""+args[0]+"\", "
  idHash              := "\"IDHash\":\""+args[1]+"\", "
  transactionType     := "\"TransactionType\":\""+args[2]+"\", "
  amount              := "\"Amount\":\""+args[3]+"\", "
  timestamp           := "\"Timestamp\":\""+args[4]+"\", "
  location            := "\"Location\":\""+args[5]+"\", "
  status              := "\"Status\": 1"

  device_json := "{"+deviceFingerPrint+idHash+transactionType+amount+timestamp+location+status+"}"
  // fmt.Println("record_transaction device_json " + device_json)
  // Convert the JSON defined above into a vehicle object for go
  err := json.Unmarshal([]byte(device_json), &d)

  // fmt.Println("record_transaction json.Unmarshal err " + err)

  if err != nil { return nil, errors.New("Invalid JSON object") }

  transaction, err := t.transact(stub, d)

  // fmt.Println("record_transaction transaction " + transaction)

  if transaction != true { 
    fmt.Printf("RECORD_TRANSACTION: Error converting transaction record: %s", transaction);
  }

  return nil, nil
}

func (t *SimpleChaincode) transact(stub shim.ChaincodeStubInterface, d DeviceTransaction) (bool, error) {

  bytes, err := json.Marshal(d)

  if err != nil { fmt.Printf("SAVE_CHANGES: Error converting device record: %s", err); return false, errors.New("Error converting device record") }

  err = stub.PutState(d.DeviceFingerPrint, bytes)

  if err != nil { fmt.Printf("SAVE_CHANGES: Error storing device record: %s", err); return false, errors.New("Error storing device record") }

  return true, nil
}


// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var key, value string
  var err error
  fmt.Println("running write()")

  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
  }

  key = args[0]
  value = args[1]
  err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
  if err != nil {
    return nil, err
  }
  return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("invoke is running " + function)
  fmt.Println("invoke why you running " + function)

  // else if function == "write" {
  //   return t.write(stub, args)

  //   // pass identity object as args
  // }

  // Handle different functions
  if function == "init" {
    return t.Init(stub, "init", args)
  } else if function == "record_transaction" {
    fmt.Println("calling record_transaction")
    // fmt.Println("args are " + args)
    return t.record_transaction(stub, args)    
  } else if function == "create_identity" {
    fmt.Println("calling create_identity")
    // fmt.Println("args are " + args)
    return t.create_identity(stub, args)
  } else {
    fmt.Println("invoke did not find func: " + function)
  }


  return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("query is running " + function)

  // Handle different functions
  if function == "read" { //read a variable
    return t.read(stub, args)
  }
  fmt.Println("query did not find func: " + function)

  return nil, errors.New("Received unknown function query: " + function)
}


// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var key, jsonResp string
  var err error

  if len(args) != 1 {
    return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
  }

  key = args[0]
  valAsbytes, err := stub.GetState(key)
  if err != nil {
    jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
    return nil, errors.New(jsonResp)
  }

  return valAsbytes, nil
}