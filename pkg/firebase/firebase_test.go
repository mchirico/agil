package firebase

import (
	"context"
	"fmt"
	gofirebase "github.com/mchirico/go-firebase/pkg/gofirebase"
	util "github.com/mchirico/go-firebase/pkg/utils"
	"testing"
)

func TestReadWrite_Firebase(t *testing.T) {

	credentials, err := FindCredentials()
	if err != nil {
		t.Logf("Not able to use a credentials file for testing.")
		return
	}

	//StorageBucket := os.Getenv("FIREBASE_BUCKET")
	StorageBucket := "septapig.appspot.com"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	number := 5
	doc := make(map[string]interface{})
	doc["application"] = "FirebaseGo"
	doc["function"] = "TestAuthenticate"
	doc["test"] = "This is example text..."
	doc["random"] = number

	fb := &gofirebase.FB{Credentials: credentials, StorageBucket: StorageBucket}
	fb.CreateApp(ctx)
	fb.WriteMap(ctx, doc, "testAgil", "AgilV0.0.1")

	resultFind, err := fb.Find(ctx, "testAgil", "function", "==", "TestAuthenticate")
	if resultFind["test"] != "This is example text..." {
		t.Fatalf("Find not working")
	}

	dsnap, _ := fb.ReadMap(ctx, "testAgil", "AgilV0.0.1")
	result := dsnap.Data()

	fmt.Printf("Document data: %v %v\n", result["random"].(int64), number)
	if result["random"].(int64) != 5 {
		t.Fatalf("Didn't return correct value\n")
	}

	util.CreateDir(".slop")
	data := []byte("ABC€")

	util.Write(".slop/junk.txt", data, 0600)
	fb.Bucket.Upload(ctx, ".slop/junk.txt")
	util.RmDir(".slop")
	err = fb.Bucket.DeleteFile(ctx, ".slop/junk.txt")

	if err != nil {
		t.Logf("Problem with buckets")
	}

}

func TestLocateFile(t *testing.T) {
	directories := []string{"/credentials",
		"../../credentials",
		"../credentials",
		"../../../credentials",
		"/etc/credentials",
	}
	locations := []string{}
	file := "septapig-firebase-adminsdk.json"
	for _, d := range directories {
		locations = append(locations, d+"/"+file)
	}

	loc, err := LocateFile(locations)
	if err != nil {
		t.Logf("Location: %s\n", loc)

	}
}

func TestFindCredentials(t *testing.T) {
	loc, err := FindCredentials()
	t.Logf("loc: %s, err: %v\n", loc, err)

	_, err = FindCredentials("a", "b", "c", "d", "e", "f")
	if err == nil {
		t.Fatalf("err should have value")
	}
}
