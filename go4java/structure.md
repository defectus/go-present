$GOPATH 
    |
    - github.com 
        |
        - you company name 
            |
            - project name
                |
                - shared/ optional packages that are shareable (structs, interfaces)
                - cmd/ 
                  |
                  - main.go - containt the main function, package name is main
                - foo/ other application specific source files 
                - bar/ more application specific files 
                - baz.go
                - baz_test.go your test files live here too

                                          