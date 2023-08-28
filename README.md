# fairOs-dfs-clients

This repo has the swagger client implementations for the fairOS dfs API.

Different branches contain different implementations for different languages.

[GO](https://github.com/fairDataSociety/fairOS-dfs-clients/tree/go)
[js](TODO)
[ts](TODO)
[python](TODO)

## Usage

You can use the branches as a boilerplate or generate your own client using the swagger file.

```
// Replace https://generator.swagger.io/api/gen/clients/<Language> at the end of this command
curl -X POST -H "content-type:application/json" -d '{"swaggerUrl":"https://raw.githubusercontent.com/fairDataSociety/fairOS-dfs-clients/main/swagger.yaml"}' https://generator.swagger.io/api/gen/clients/go
```

Generator has support for these languages
```
["ada","akka-scala","android","apex","bash","clojure","cpprest","csharp","csharp-dotnet2","cwiki","dart","dart-jaguar","dynamic-html","eiffel","elixir","elm","erlang-client","flash","go","groovy","haskell-http-client","html","html2","java","javascript","javascript-closure-angular","jaxrs-cxf-client","jmeter","kotlin","lua","objc","perl","php","powershell","python","qt5cpp","r","ruby","rust","scala","scalaz","swagger","swagger-yaml","swift","swift3","swift4","swift5","tizen","typescript-angular","typescript-angularjs","typescript-aurelia","typescript-fetch","typescript-inversify","typescript-jquery","typescript-node","ue4cpp"]
```

## Contributing

If you want to add a new language, please create a new branch and open a PR.