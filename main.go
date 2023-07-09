package main

import (
  "context"
  "fmt"
  "net/http"

  "github.com/qorpress/go-wordpress"
)

func main() {

  tp := wordpress.BasicAuthTransport{
    Username: USER,
    Password: PASSWORD,
  }

  // create wp-api client
  client, _ := wordpress.NewClient(API_BASE_URL, tp.Client())

  ctx := context.Background()

  // for eg, to get current user (GET /users/me)
  currentUser, resp, _ := client.Users.Me(ctx, nil)
  if resp != nil && resp.StatusCode != http.StatusOK {
    // handle error
  }

  // Or you can use your own structs (for custom endpoints, for example)
  // Below is the equivalent of `client.Posts.Get(ctx, 100, nil)`
  var obj MyCustomPostStruct
  resp, err := client.Get(ctx, "/posts/100", nil, &obj)
  // ...

  fmt.Printf("Current user %+v", currentUser)
}

