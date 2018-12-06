- get all tabs 
    ##### Request example:
            GET https://localhost:8080/api/v1/tabs?enabled=true
                enabled=true    - only enabled tabs
                enabled=false   - all tabs (with disabled)
    ##### Successfull response example:
            [
               {
                  "id":1,
                  "title":"HOME",
                  "user_type_visible":
                  [
                     "all"
                  ],
                  "index":1,
                  "enabled":true,
                  "tab_type":
                  {
                     "id":1,
                     "type":"home"
                  }
               },
               {
                  "id":2,
                  "title":"БЛОГ",
                  "user_type_visible":
                  [
                     "all"
                  ],
                  "index":2,
                  "enabled":true,
                  "tab_type":
                  {
                     "id":2,
                     "type":"blog"
                  }
               }
            ]

- register (create) user
    ##### Request example:
            POST https://localhost:8080/api/v1/user
            POST data:
                {
                	"nickname": "pupkin",
                	"password": "12345678",
                	"email":"pupkin@example.com",
                	"phone":"79163777604",
                	"name":"Василий",
                	"lastname":"Пупкин"
                }
    ##### Successfull response example:  
        StatusOK
        empty response    
              
- login (create session)
    ##### Request example:
        POST https://localhost:8080/api/v1/session
        POST data:
            {
        	    "login": "test",
        	    "password": "1234"
            }
    ##### Successfull response example:  
         {
            "id":4,
            "email":"test@example.com",
            "phone":"",
            "nickname":"test",
            "name":"test",
            "last_name":"",
            "role":
            {
               "id":2,
               "role":"client"
            }
         }   
                            
- logout (destroy session)
    ##### Request example:
        DELETE https://localhost:8080/api/v1/session
    ##### Successfull response example:  
         StatusOK
         empty response 
                            
- set tabs state            
     ##### Request example:
        PUT https://localhost:8080/api/v1/tabs
        PUT Body:
                [
                   {
                      "id":1,
                      "title":"HOME",
                      "user_type_visible":
                      [
                         "all"
                      ],
                      "index":1,
                      "enabled":true,
                      "tab_type":
                      {
                         "id":1,
                         "type":"home"
                      }
                   },
                   {
                      "id":2,
                      "title":"БЛОГ",
                      "user_type_visible":
                      [
                         "all"
                      ],
                      "index":2,
                      "enabled":true,
                      "tab_type":
                      {
                         "id":2,
                         "type":"blog"
                      }
                   }
                ]
     ##### Successfull response example:    
            StatusOK
            empty response