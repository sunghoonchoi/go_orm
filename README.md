# sunghoon_choi__simple_orm


Title
=============================
Simple ORM module for golang environment




my core mission
=============================
1. easy to use
    -> don't need many explanations
    -> clear to know how to use this library
2. maintance
    -> support other data types later
    -> support other database (not sqlite3) later
    -> consideration of security




dev methology
=============================
1. test driven development using test case
2. small unit test based development create, update, select, delete
3. but I didn't make not all of _test.go files for the test
4. i wrote detail use instructions in the main.go (sample)




how to use this orm module
=============================
1. create your struct

2. regist to orm engine
    -> it will return db object (db is object for orm)
    -> user need this db object to use ddl ro queries
    
3. how to use this module
    
    ```
    1.define your own struct
        type your_struct struct {
            int, string, byte[]
        }
    2.create table
        RegisterInOrmEngine <- let this method know the primary key index
    3.insert
        InsertToTable
    4.update
        UpdateRows
           + []SearchConditon
           + []TargetColumn 
    5.delete
        DeleteRows
           + []SearchConditon
    6.search
        SearchRows
           + []SearchConditon
    ```
    



Test
=============================
<IF GO INSTALLED>

```
1. move to source folder
2. go run .
```
OR
    
<IF NOT>

```
run compiled file "orm_sample"
```


    
    
Developer note
=============================
1. basic operations are working now
2. remains
    1. search and return result -> complete
    2. error define and handling -> basic handling complete
    3. review the usage convenience -> complete
    4. prevent misuse and test exception case -> remain
    5. error code, message define and mapping -> remain
    6. search conditon with time -> remain
3. big change of design
    1. struct in struct model
        * I was preparing every struct which want to be store in the db have orm_engin inside but it was changed because of the go can't access to another packages struct
    2. interface model
        * golang don't have file for each class so the code is not seperated well so I did't use this model
    3. controller model
        * orm controller (crud) do the request of the struct.
    
4. hard problems during making this
    1. making common function for search
        -> deep understand of interface and get the element of pointer by Elem() method
    2. runtime variable declare and mapping to sqlite function
        -> type check and dynamically append to interface slice

