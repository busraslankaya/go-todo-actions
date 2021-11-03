# 📝 TO-DO LIST APP

This is a to-do list application developed with Golang.

Local connection is provided with the :8000 port specified in the main.go file. There are five functions named indexHandler, TodosHandler, GetTodoListHandler, AddTodoListHandler, NewHandler in the app.go file. 
A test function is executing for the AddTodoListHandler function in the file named app_test.go.

I also created docker image with Dockerfile file.
I sent my yml files in github/workflows folder to actions and built it.


*BackEnd*

  /todo: add a new TODO list / POST
  
  /todo: add list via index.html / POST
  
  /todo: Print the entire list of TODOs / GET
  
  /todo/{id:[0-9]+}: Display the TODO list corresponding to id / GET
  
  /todo: update TODO list / PUT
  
  /todo/{id:[0-9]+}: Created so that TODOLIST Content can be edited / PUT
  
  
  
Frontend

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
