package main

// global variable declaration
var c, java, python bool = true, false, true

// Pi global constant variables declaration
const Pi = 3.14

func main() {
	// go variable types
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32
	// float32 float64
	// complex64 complex128
	onenewSlice := make([]int, 5)
	fmt.Println(onenewSlice, len(onenewSlice), cap(onenewSlice))

	//lets create slices of slices

	slicesOfSlices := [][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}

	for i := 0; i < len(slicesOfSlices); i++ {
		fmt.Println(strings.Join(slicesOfSlices[i], " "))
	}
	println(Pi)
	// type conversions
	bhaskar := "hero"
	rank := 1
	height := 5.11

	println(bhaskar, float32(rank), int(height))

	// this is a simple example demonstrating the declaration of variable in var
	var x int = 10
	println(x, c, java, python)
	// this is an example demonstarting the declaring and initializing of variables
	a, b := swappingOfNumbers(10, 20)
	println(a, b)
	//variables with initilizers so that we can omit the type
	var i, j, k = 1, 2, 3
	println(i, j, k)

	// Now lets see how to create a server in golang 
	// before jumping straight into it lets see what is the underlying concept behind creating http servers in general
	// there will be a mulitplexer/router that sits infront of server and routes the incoming requests to specific routes registered to the multilpexer/router
	// You can consider express as a multiplexer for NodeJS HTTP server
	// In golang there is a struct called as ServeMux in http package which is the inbuilt implementation of multiplexer/router

	mux := http.NewServeMux{}

	// now define the functions which will handle the routes

	function firstRouteHandler(w http.Response,r *http.RequestHandler){
	// handle the request and send the response
	}

	function secondRouteHandler(...){
	// handle the request	
	}

	function secondPostRouterHandler(...){
		if r.Method != http.MethodPost{
			fmt.Println("err")
			return
		}
	}

	// by default mux register handle get request only
	mux.HandleFunc("/first" , firstRouteHandler)
	mux.HandleFunc("/second" , secondRouteHandler)
        // in order to handle Post request also
	mux.HandleFunc("/second" , secondPostRouteHandler)
	// create the server and run it

	err := http.ListenAndServe(PORT , mux)
	if(err != nil){
		fmt.Prrintln("error starting the server")
	}


	// now we will see how to use middlewares in golang
        // middlewares are the functions which are executed right before the request is handled or very next the response is sent
	// mainly used for authentications, tokens , cookies things like that 

	function middleWareWrapping(next Http.Handler) http.Handler{
		return http.HandlerFunc(func(w ,r){
			// implement middleware logic for handling request
                           next.ServeHTTP(w,r)
			
		})
		
	}

	function thirdRouteHandler(...){
	}

	wrappedThirdRouteHandler := middleWareWrapping(thirdRouteHandler)

	mux.HandleFunc("/third" , wrappedThirdRouteHandler)
	
}

// function demonstrating how the return works in go
func additionOfNumbers(x int, y int) int {
	return x + y
}

// function demonstrating how the return can handle two or more than two values
func swappingOfNumbers(x int, y int) (int, int) {
	return y, x
}

// function demonstrating how the naked return works in go
func nakedReturn(x int, y int) (a int, b int) {
	a = y
	b = x
	return
}
