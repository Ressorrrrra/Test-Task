type Order struct{
ID int 
Items []*Item 
Cost int 
OrderedAt int64
}


type Item struct{
ID int 
Name string 
Description string
Price int 
}