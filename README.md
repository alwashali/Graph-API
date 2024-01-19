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

```
./main -post -endpoint https://graph.microsoft.com/v1.0/security/runHuntingQuery -data '{"Query": "EmailEvents | take 1 "}' | jq
```

Replace the endpoint URL with your desired Graph API endpoint to test other functionalities.

---

## Use curl 


**Set Credentilas**

```
CLIENT_ID="your-client-id"
CLIENT_SECRET="your-client-secret"
TENANT_ID="your-tenant-id"
```

**Request Token**

```
response=$(curl -s -X POST "https://login.microsoftonline.com/$TENANT_ID/oauth2/v2.0/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "client_id=$CLIENT_ID" \
     -d "scope=https://graph.microsoft.com/.default" \
     -d "client_secret=$CLIENT_SECRET" \
     -d "grant_type=client_credentials");token=$(echo $response | jq -r '.access_token')

```

**Test Endpoints**

```
curl -H "Authorization: Bearer $token" \
     -H "Content-Type: application/json" \
     https://graph.microsoft.com/v1.0/users
```

---


## Office365 Management Activity API

**Request Token**

scope=https://manage.office.com/.default 

```
response=$(curl -s -X POST "https://login.microsoftonline.com/$TENANT_ID/oauth2/v2.0/token" \
	     -H "Content-Type: application/x-www-form-urlencoded" \
	     -d "client_id=$CLIENT_ID" \
	     -d "scope=https://manage.office.com/.default" \
	     -d "client_secret=$CLIENT_SECRET" \
	     -d "grant_type=client_credentials" \
	     -d "redirect_uri=http://localhost");token=$(echo $response | jq -r '.access_token')
```


**Start Subscription**

```
curl -XPOST -H "Authorization: Bearer $token" \
     -H "Content-Type: application/json" \
		 -d ''\
     'https://manage.office.com/api/v1.0/$TENANT_ID/activity/feed/subscriptions/start?contentType=Audit.General'
```


**List Content**


```
curl -H "Authorization: Bearer $token" \
     -H "Content-Type: application/json" \
    'https://manage.office.com/api/v1.0/$TENANT_ID/activity/feed/subscriptions/content?contentType=Audit.General' 
```

Response

```
[
     {
         "contentUri": "Content URI",
         "contentId": "ContentID",
         "contentType": "Audit.General",
         "contentCreated": "2024-01-19T09:09:19.175Z",
         "contentExpiration": "2024-01-25T15:11:08.116Z"
       },
     ....
]
```


**Fetch Logs**

```
curl -H "Authorization: Bearer $token" \
     -H "Content-Type: application/json" \
    'Content URI' 
```


