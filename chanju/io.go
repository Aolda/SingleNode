package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// go run io.go -action "insert or delete" -name "파일명" -> 추후에 update나 db 같은 내용도 추가 가능
func main() {
	action := flag.String("action", "", "insert or another")
	name := flag.String("name", "", "what file")

	flag.Parse()

	if *action == "insert"{
	insertJs(*name)
	}else if *action == "delete"{
	deletJs(*name)
	}else if *action == "update"{
	updateJs(*name)
	}else{
		fmt.Println("등록ㄴㄴ\n");
	}
}

// io.go와 같은 폴더에 위치해야지 넣을 수 있음
// 지정된 db에 저장 -> 추후 수정 가능
func insertJs(name string){
	
	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := filepath.Join(cwd, name)
	filePath = filepath.Base(string(filePath))

	fmt.Println(filePath)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}

	collection := client.Database("mydb").Collection("files")
	_, err = collection.InsertOne(context.TODO(), bson.M{"name": name, "data":file})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("잘 저장됨")
	}

	err = client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
}

func deletJs(name string ){
	 client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
 
	 collection := client.Database("mydb").Collection("files")
	 
	 filter := bson.M{"name": name}

	 deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 fmt.Printf(" %v 지움\n", deleteResult.DeletedCount)
 
	 err = client.Disconnect(context.TODO())
	 if err != nil {
		 fmt.Println(err)
	 }
}

func updateJs(name string){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		  fmt.Println(err)
		  return
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Connect(ctx)
    if err != nil {
        fmt.Println(err)
        return
    }

	collection := client.Database("mydb").Collection("files")

	filter := bson.M{"name": name}
    var result bson.M
    err = collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            fmt.Println("Document not found.")
            return
        }
        fmt.Println(err)
        return
    }

	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := filepath.Join(cwd, name)
	filePath = filepath.Base(string(filePath))

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	filter = bson.M{"name": name}
	update := bson.M{"$set": bson.M{"data": file}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return
	}
  
	fmt.Println("잘 갱신됨")
}