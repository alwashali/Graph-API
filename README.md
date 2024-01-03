# Graph-API
Test Microsoft Graph API Endpoints 


**Registering Your Application**

To use the Graph-API Explorer, first register your application in Azure Active Directory.
update the necessary credentials in the config file (Client ID, Client Secret, Tenant ID) and set the required permissions for your application.

**Installatio**

Clone this repository and build the project:
```
git clone https://github.com/alwashali/Graph-API.git
cd Graph-API
go build main.go
```


**Examples**
The following commands can be used to test different Graph API endpoints:


```
./main -endpoint https://graph.microsoft.com/v1.0/identityProtection/riskyUsers | jq
```

```
./main -endpoint https://graph.microsoft.com/v1.0/users | jq
```

Replace the endpoint URL with your desired Graph API endpoint to test other functionalities.
