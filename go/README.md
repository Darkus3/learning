# Go
After watching rwxrob streams, I figure out that go might be intereting
to learn and play with. It seems powerfull tool to do testing and
scripting.



## Agenda
https://go.dev/doc/tutorial/create-module
- ~~Create a module -- Write a small module with functions you can call from another module.~~
- ~~Call your code from another module -- Import and use your new module.~~
- ~~Return and handle an error -- Add simple error handling.~~
- ~~Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).~~
- ~~Return greetings for multiple people -- Store key/value pairs in a map.~~
- Add a test -- Use Go's built-in unit testing features to test your code.
- Compile and install the application -- Compile and install your code locally.

### DAY 0
Build simple module and discover basics of go lang
dependencies:
- fmt (standart in & out): Println, Scanf, Sprintf...
- errors
- log


### DAY 1
- Install vim-go plugin to extend functionality: auto-completion etc...
  - Uprade vim to 8.2, gcc, make & ncurses dependencies
- Modify greet module to generate random message 
- Add multi names handling with "Slice", that are dynamic tables.
- Use Hash Map to associate greet message with names.
- Write *Test case* for my module. Go make it easy to deploy unit
  testing with internal tools and stdlib through module *testing*.
  - TestHelloName := Test return value if a match with original name is
    found
  - TestHelloEmpty := Test is error handling works, should fail if Hello
    return empty string and no error
- Build module and add go bin dir to $PATH. You can now run your module
  as a shell cmd.
