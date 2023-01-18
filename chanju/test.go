package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/robertkrimen/otto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogEvent struct {
    Message string
}

func main() {
   //storgefile()
   fineFileAndExecute()
   //deleteMongo()
}

// func ConnectEth(){
// 	client, err := ethclient.Dial("wss://mainnet.infura.io/ws")
//     if err != nil {
//         log.Fatal(err)
//     }

//     contractAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
//     query := bind.NewBoundContract(contractAddress, abi, client)

//     // Watch for events
//     ctx := context.Background()
//     events := make(chan LogEvent)
//     query.WatchLogEvent(ctx, events)

//     for event := range events {
//         log.Printf("Event received: %+v", event)
//     }
// }

func fineFileAndExecute(){
	 // Connect to MongoDB
	 client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	 if err != nil {
		 log.Fatal(err)
	 }
	 ctx := context.Background()
	 err = client.Connect(ctx)
	 if err != nil {
		 log.Fatal(err)
	 }
 
	 // Get a handle to the files collection
	 files := client.Database("mydb").Collection("files")
	 // Define the filter to find the file
	 filter := bson.M{"name": "file.js"}
	 // Find the file
	 var file bson.M
	 err = files.FindOne(ctx, filter).Decode(&file)
	 if err != nil {
		 log.Fatal(err)
	 }
	 // check if the file is not empty
	 if len(file) == 0 {
		 log.Fatal("File not found")
	 }
	data := file["data"].(primitive.Binary).Data

	  // Create a new JavaScript VM
	vm := otto.New()
	  // Run the JavaScript file
	result, err := vm.Run(string(data))
	if err != nil {
		log.Fatal(err)
	}
	dateString, _ := result.ToString()
	fmt.Println("The current date is:", dateString)
}







func deleteMongo(){
	 // Connect to MongoDB
	 client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
 
	 // Get a handle to the collection
	 collection := client.Database("mydb").Collection("files")
 
	 // Delete all documents in the collection
	 _, err = collection.DeleteMany(context.TODO(), bson.M{})
	 if err != nil {
		 fmt.Println(err)
	 } else {
		 fmt.Println("All documents successfully deleted from the collection!")
	 }
 
	 // Close the connection to the database
	 err = client.Disconnect(context.TODO())
	 if err != nil {
		 fmt.Println(err)
	 }
	 
}

func storgefile(){
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the local .js file
	file, err := ioutil.ReadFile("./file.js")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Insert the file into MongoDB
	collection := client.Database("mydb").Collection("files")
	_, err = collection.InsertOne(context.TODO(), bson.M{"name": "file.js", "data":file})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File successfully saved to MongoDB!")
	}

	// Close the connection to the database
	err = client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
}