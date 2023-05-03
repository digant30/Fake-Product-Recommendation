package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// Product :  Define the product structure, with 8 properties.  Structure tags are used by encoding/json library
type Product struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Brand  string `json:"brand"`
	Colour string `json:"colour"`
	Price  string `json:"price"`
	Size   string `json:"size"`
	Owner  string `json:"owner"`
	SelId  string `json:"sellerid"`
}

// Seller :  Define the seller structure, with 6 properties.  Structure tags are used by encoding/json library
type Seller struct {
	SelId   string `json:"sellerid"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
	MgrId   string `json:"mgrid"`
	Address string `json:"address"`
	Brand   string `json:"brand"`
}

// Consumer :  Define the product structure, with 8 properties.  Structure tags are used by encoding/json library
type Consumer struct {
	ProdId   string `json:"prodid"`
	ProdName string `json:"prodname"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Contact  string `json:"contact"`
	Mail     string `json:"mail"`
	SelId    string `json:"sellerid"`
	SelName  string `json:"sellername"`
}
type Review struct {
	ProdName string `json:"prodname"`
	SelName  string `json:"sellername"`
	Review   string `json:"review"`
}

// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("smartContract_cc")

// Invoke :  Method for INVOKING smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args))

	switch function {
	case "queryProduct":
		return s.queryProduct(APIstub, args)
	case "querySeller":
		return s.querySeller(APIstub, args)
	case "queryConsumer":
		return s.queryConsumer(APIstub, args)
	case "initLedger":
		return s.initLedger(APIstub)
	case "createProduct":
		return s.createProduct(APIstub, args)
	case "queryAllProducts":
		return s.queryAllProducts(APIstub)
	case "createSeller":
		return s.createSeller(APIstub, args)
	case "queryAllSellers":
		return s.queryAllSellers(APIstub)
	case "createConsumer":
		return s.createConsumer(APIstub, args)
	case "queryAllConsumers":
		return s.queryAllConsumers(APIstub)
	case "changeProductOwner":
		return s.changeProductOwner(APIstub, args)
	case "getHistoryForAsset":
		return s.getHistoryForAsset(APIstub, args)
	case "queryProductByOwner":
		return s.queryProductByOwner(APIstub, args)
	case "queryConsumerContact":
		return s.queryConsumerContact(APIstub, args)
	case "authenticate":
		return s.authenticate(APIstub, args)
	case "addreview":
		return s.addreview(APIstub, args)
	case "queryreview":
		return s.queryreview(APIstub, args)
	case "restictedMethod":
		return s.restictedMethod(APIstub, args)
	case "test":
		return s.test(APIstub, args)
	case "queryPrivateDataHash":
		return s.queryPrivateDataHash(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}

	// return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(productAsBytes)
}

func (s *SmartContract) querySeller(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sellerAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(sellerAsBytes)
}

func (s *SmartContract) queryConsumer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	consumerAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(consumerAsBytes)
}

func (s *SmartContract) test(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(productAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	products := []Product{
		Product{Id: "A0000001", Name: "Runtastic", Brand: "Adidas", Colour: "blue", Price: "5000", Size: "M", Owner: "Adidas Porur Chennai", SelId: "S28432"},
	}

	sellers := []Seller{
		Seller{SelId: "S28432", Name: "JP Stores", Manager: "JP Naidu", MgrId: "10293745", Address: "Vadapalani", Brand: "Adidas"},
	}

	consumers := []Consumer{
		Consumer{ProdId: "A0000001", ProdName: "Adidas Runtastic", Name: "Shilpa", Address: "Guindy", Contact: "9191923456", Mail: "shilpa@gmail.com", SelId: "S28432", SelName: "JP Stores"},
	}

	review := []Review{
		Review{ProdName: "Adidas Runtastic", SelName: "JP Stores", Review: "5"},
	}

	i := 0
	for i < len(products) {
		productAsBytes, _ := json.Marshal(products[i])
		APIstub.PutState("ITEM"+strconv.Itoa(i), productAsBytes)
		i = i + 1
	}

	j := 0
	for j < len(sellers) {
		sellersAsBytes, _ := json.Marshal(sellers[j])
		APIstub.PutState("SELLER"+strconv.Itoa(j), sellersAsBytes)
		j = j + 1
	}

	k := 0
	for k < len(consumers) {
		consumerAsBytes, _ := json.Marshal(consumers[k])
		APIstub.PutState("CONSUMER"+strconv.Itoa(k), consumerAsBytes)
		k = k + 1
	}

	l := 0
	for l < len(review) {
		reviewAsBytes, _ := json.Marshal(review[l])
		APIstub.PutState("REVIEW"+strconv.Itoa(l), reviewAsBytes)
		l = l + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	var product = Product{Id: args[1], Name: args[2], Brand: args[3], Colour: args[4], Price: args[5], Size: args[6], Owner: args[7], SelId: args[8]}

	productAsBytes, _ := json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	indexName := "owner~key"
	colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{product.Owner, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)

	return shim.Success(productAsBytes)
}

func (s *SmartContract) createSeller(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var seller = Seller{SelId: args[1], Name: args[2], Manager: args[3], MgrId: args[4], Address: args[5], Brand: args[6]}

	sellerAsBytes, _ := json.Marshal(seller)
	APIstub.PutState(args[0], sellerAsBytes)

	indexName := "name~key"
	colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{seller.Name, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)

	return shim.Success(sellerAsBytes)
}

func (s *SmartContract) createConsumer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	var consumer = Consumer{ProdId: args[1], ProdName: args[2], Name: args[3], Address: args[4], Contact: args[5], Mail: args[6], SelId: args[7], SelName: args[8]}

	consumerAsBytes, _ := json.Marshal(consumer)
	APIstub.PutState(args[0], consumerAsBytes)

	indexName := "ProdId~key"
	colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{consumer.ProdId, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)

	indexName1 := "contact~key"
	colorNameIndexKey1, err1 := APIstub.CreateCompositeKey(indexName1, []string{consumer.Contact, args[0]})
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	value1 := []byte{0x00}
	APIstub.PutState(colorNameIndexKey1, value1)

	return shim.Success(consumerAsBytes)
}

func (s *SmartContract) addreview(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var review = Review{ProdName: args[1], SelName: args[2], Review: args[3]}

	productAsBytes, _ := json.Marshal(review)
	APIstub.PutState(args[0], productAsBytes)

	indexName := "SelName~key"
	colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{review.SelName, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)

	indexName2 := "ProdName~key"
	colorNameIndexKey2, err2 := APIstub.CreateCompositeKey(indexName2, []string{review.ProdName, args[0]})
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	value2 := []byte{0x00}
	APIstub.PutState(colorNameIndexKey2, value2)

	return shim.Success(productAsBytes)
}

func (S *SmartContract) queryreview(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	prodName := args[0]

	prodAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("ProdName~key", []string{prodName})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer prodAndIdResultIterator.Close()

	var i int
	var id string

	var products []byte
	bArrayMemberAlreadyWritten := false

	products = append([]byte("["))

	for i = 0; prodAndIdResultIterator.HasNext(); i++ {
		responseRange, err := prodAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		id = compositeKeyParts[1]
		assetAsBytes, err := APIstub.GetState(id)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), assetAsBytes...)
			products = append(products, newBytes...)

		} else {
			products = append(products, assetAsBytes...)
		}

		fmt.Printf("Found a asset for index : %s asset id : %s %s", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	products = append(products, []byte("]")...)

	return shim.Success(products)
}

func (S *SmartContract) queryProductByOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	owner := args[0]

	ownerAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("owner~key", []string{owner})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer ownerAndIdResultIterator.Close()

	var i int
	var id string

	var products []byte
	bArrayMemberAlreadyWritten := false

	products = append([]byte("["))

	for i = 0; ownerAndIdResultIterator.HasNext(); i++ {
		responseRange, err := ownerAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		id = compositeKeyParts[1]
		assetAsBytes, err := APIstub.GetState(id)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), assetAsBytes...)
			products = append(products, newBytes...)

		} else {
			products = append(products, assetAsBytes...)
		}

		fmt.Printf("Found a asset for index : %s asset id : %s %s", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	products = append(products, []byte("]")...)

	return shim.Success(products)
}

func (S *SmartContract) authenticate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	prodId := args[0]

	ownerAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("ProdId~key", []string{prodId})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer ownerAndIdResultIterator.Close()

	var i int
	var id string

	var products []byte
	bArrayMemberAlreadyWritten := false

	products = append([]byte("["))

	for i = 0; ownerAndIdResultIterator.HasNext(); i++ {
		responseRange, err := ownerAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		id = compositeKeyParts[1]
		assetAsBytes, err := APIstub.GetState(id)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), assetAsBytes...)
			products = append(products, newBytes...)

		} else {
			products = append(products, assetAsBytes...)
		}

		fmt.Printf("Found a asset for index : %s asset id : %s %s", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	products = append(products, []byte("]")...)

	return shim.Success(products)
}

func (S *SmartContract) queryConsumerContact(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	contact := args[0]

	ownerAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("contact~key", []string{contact})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer ownerAndIdResultIterator.Close()

	var i int
	var id string

	var products []byte
	bArrayMemberAlreadyWritten := false

	products = append([]byte("["))

	for i = 0; ownerAndIdResultIterator.HasNext(); i++ {
		responseRange, err := ownerAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		id = compositeKeyParts[1]
		assetAsBytes, err := APIstub.GetState(id)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), assetAsBytes...)
			products = append(products, newBytes...)

		} else {
			products = append(products, assetAsBytes...)
		}

		fmt.Printf("Found a asset for index : %s asset id : %s %s", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	products = append(products, []byte("]")...)

	return shim.Success(products)
}

func (s *SmartContract) queryAllProducts(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "ITEM0"
	endKey := "ITEM999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllProducts:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllSellers(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "SELLER0"
	endKey := "SELLER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllSellers:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllConsumers(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CONSUMER0"
	endKey := "CONSUMER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllConsumers:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) restictedMethod(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	val, ok, err := cid.GetAttributeValue(APIstub, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		shim.Error("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		shim.Error("Client identity doesnot posses the attribute")
	}
	// Do something with the value of 'val'
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return shim.Error("Only user with role as APPROVER have access this method!")
	} else {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		productAsBytes, _ := APIstub.GetState(args[0])
		return shim.Success(productAsBytes)
	}

}

func (s *SmartContract) changeProductOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	product := Product{}

	json.Unmarshal(productAsBytes, &product)
	product.Owner = args[1]

	productAsBytes, _ = json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	indexName := "owner~key"
	colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{product.Owner, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	APIstub.PutState(colorNameIndexKey, value)

	return shim.Success(productAsBytes)
}

func (t *SmartContract) getHistoryForAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productName := args[0]

	resultsIterator, err := stub.GetHistoryForKey(productName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForAsset returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryPrivateDataHash(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	productAsBytes, _ := APIstub.GetPrivateDataHash(args[0], args[1])
	return shim.Success(productAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
