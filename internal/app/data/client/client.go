type Order struct{
ID int 
Items []*Item 
Cost int 
OrderedAt int64
}

type Client struct{
ID int 
Username string 
Email string
Password string
Orders []*Order
Sort int
}

type Item struct{
ID int 
Name string 
Description string
Price int 
}