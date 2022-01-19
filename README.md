# shipment app

###To run the application:
``` 
go run main.go
```
###Endpoints:

- http://localhost:8080/v1/api/shipments (get all shipments)
- http://localhost:8080/v1/api/shipments/:id (get shipment by id)
- http://localhost:8080/v1/api/shipments (post shipment)

Post request body example:

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
