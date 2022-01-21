# shipment app

### To run the application:
``` 
go run main.go
```
### Endpoints:

- Get all shipments (GET):
  http://localhost:8080/v1/api/shipments 
- Get shipment by id (GET):
  http://localhost:8080/v1/api/shipments/:id
- Post shipment (POST):
  http://localhost:8080/v1/api/shipments 

Post request body example (for postman):
``` 
  {
   "senderName": "Nicolas",
   "senderEmail": "Nicolas.Cage@gmail.com",
   "senderAddress": "California",
   "senderCountryCode": "US",
   "recipientName": "Bill",
   "recipientEmail": "Bill.Murray@gmail.com",
   "recipientAddress": "Evanston",
   "recipientCountryCode": "US",
   "weight": 15
   }
```

### SQLite for storage
