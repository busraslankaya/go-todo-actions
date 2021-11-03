# 📝 WEB BASED TO-DO LIST APPLICATION

This is a to-do list application developed with Golang.

Local connection is provided with the :8000 port specified in the main.go file. 
A test function is executing for the TestIndexGet and TestPostTodo functions in the file named app_test.go. But I did not get successful results in this test file.
![image](https://user-images.githubusercontent.com/60943616/140061750-cdfa6c7b-b7e7-40b7-8558-868a34e9b8e8.png)


I also created docker image with Dockerfile file.
When i run build command
![image](https://user-images.githubusercontent.com/60943616/140062754-d55d3d29-0d31-4fec-87b9-43d6d382092e.png)
I created my docker image
![image](https://user-images.githubusercontent.com/60943616/140062829-251bfabc-9b39-4f45-890b-234d541a1b99.png)
Iran 
![image](https://user-images.githubusercontent.com/60943616/140062893-c42aaaaf-2d66-4dde-affa-a84ef26dfa9f.png)

And here is my interface.

![image](https://user-images.githubusercontent.com/60943616/140058038-ceda5aad-da04-4664-9a01-15242ef27544.png)

![image](https://user-images.githubusercontent.com/60943616/140062546-d3ccb54e-ee31-4e96-9333-86d216d882d0.png)


I sent my yml files in github/workflows folder to actions and built it.


![image](https://user-images.githubusercontent.com/60943616/140064011-c27346e8-b699-49ce-8039-6e335e0c780c.png)
![image](https://user-images.githubusercontent.com/60943616/140064135-06fc34b3-df9f-4936-b6d5-d29ddcf22fce.png)


**BackEnd**

  /todo: add a new TODO list / POST
  
  /todo: add list via index.html / POST
  
  /todo: Print the entire list of TODOs / GET
  
  /todo/{id:[0-9]+}: Display the TODO list corresponding to id / GET
  
  /todo{id:[0-9]+}: delete TODO list / DELETE
  
  /todo: update TODO list / PUT
  
  /todo/{id:[0-9]+}: Created to update the status of TODO items through Bool type / PUT
  
  /todo/{id:[0-9]+}: Created so that TODOLIST Content can be edited / PUT
  
  
  
  
**Frontend**

  /todo: print list through index.html
  
  index.html: display each item separately
  
  index.html: Marshal by API JSON id
  
  index.html: Made to be removable in the TODO item frontend
  
  index.html: Modifications to Completed items using checkboxes /Whether or not completed
  
  
// Request | POST http://localhost:8000/todo 
{

    "content": "Buy some milk" 
    
}

// Response

    {
    
    "id":1,"Content":"buy some milk","Completed":false,"created_at":"2021-11-02T22:15:23.9938856Z"
    
    }
