# fairOs-dfs-clients

This is a demo go client for the fairOS-dfs.

## How to generate client 
```
curl -X POST -H "content-type:application/json" -d '{"swaggerUrl":"https://raw.githubusercontent.com/fairDataSociety/fairOS-dfs-clients/main/swagger.yaml"}' https://generator.swagger.io/api/gen/clients/go
```

and then you get response like: `{"code":"d4b85e57-aba1-4e8c-a4b0-f0650cfb8521","link":"https://generator.swagger.io/api/gen/download/d4b85e57-aba1-4e8c-a4b0-f1650cfb8521"}`

where link points to generated library.

> Note: You can clone this repository and use it as a starting point for your own client.
> or you can generate your own go client with the above steps and put it in your repository.