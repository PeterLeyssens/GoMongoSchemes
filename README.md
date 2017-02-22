Go MongoDB Schemes example
==========================

An example implementation of polyform database entries in Golang.

Programmed using Golang 1.7.5 and MongoDB 3.4.2.

-------------------

## Disclaimer

This code depends on a database called "test", which should contain a pre-filled collection called "VehicleSchemes". The format of VehicleSchemes is as follows:

```
{
	"name": (name of the scheme),
	"fields": {
		"1": (name of field 1)
		"2": (name of field 3)
		...
	}
}
```


-------------------

## Install

Install Go and MongoDB. Make sure MongoDB is running locally using the standard settings. Create a collection "VehicleSchemes" in the database "test" and pre-fill it with a number of schemes. For example, this is the export of the VehicleSchemes collection I used for the article (without the _id fields):
```
{
  "name": "Plane",
  "fields": {
    "1": "BrandName",
    "2": "Type",
    "3": "Seats"
  }
}
{
  "name": "Spacecraft",
  "fields": {
    "1": "Name",
    "2": "Country"
  }
}
{
  "name": "Car",
  "fields": {
    "1": "BrandName",
    "2": "Type",
    "3": "Fuel",
    "4": "Doors"
  }
}
```

