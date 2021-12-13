package main

/*
-Go is not object-oriented language
-Go is not class and object  language
-Go doesn't have Private data
-Go doesn't support inheritance
-Go doesn't have abstract base types
*/

/*
-Go have Method : we can define data structure and attach method to that data structure
-Go have a concept of Package-Level-Data , so we can make our data private at package level
-So , we will focus on package oriented design instead of class oriented design, and  we can implement object-oriented concepts
-Go have Type embed mechanism , so we can implement inheritance by composition approach.
-Go support interface , go relied heavily on interface
*/

/*
-Encapsulation
 	-is all about accessing service on an object without knowing how that service is implemented
-Encapsulation in Go
	-Go has no class , no private concept
-Strategies to apply Encapsulation in Go
	-Package Oriented Design
	-Interface
-Package Oriented Design
	-First, Define our package
	-Second, Define our Data Structure , set attributes to be private(small case)
	-Third, Assign some accessor method to it

	-Conventions
		-Name Getter Method Like attribute name
        -Prefix Setter Method with Set
	-Tips
		-If you just want data-structure to hold data to retrieve it later
		-the best approach to keep his attribute public with no getter/setter method
        -treats package as the lowest level organisation unit
        -Every thing in package , should be cohesively work together
		-So, when designing  Go app , design it as series of Go packages that are interacting together
-Interface
  -the other tool, that we can use to encapsulate our services

*/

/*
-Message Passing between objects
	-sending messages to an object without knowing how this object will process the message
	-plz note message passing is different approach than Conventional Invocation
	-Conventional Invocation : in this approach the sender object call directly  the Receiver method directly and pass to it the message
	-in message passing approach , the sender  will send message to Proxy without knowing the receiver and the receiver
	 will listen to that proxy
	-by this we make lose coupling between sender and receiver

-Strategies to apply Message Passing  in Go :our goal is to abstract the way tha sender interact with receiver
	-interface
	-channels

*/

/*
-inheritance and composition

	-inheritance: behavior reuse  strategy  where  allowing  inheritance from the base type attribute and functionality
    -composition: behavior reuse  strategy  where type contain object which have desired functionality.
                  the type delegates calls to those objects to use their behaviors.

	-Inheritance is not supported in Go Language cuz it sides effects
		-Tightly couples between parent and child
		-Hard to debug and maintain in deeply tree
        -difficulty to know where  specific behavior is coming from with override capabilities
		 especially when you have deep inherited tree
		- All the functionality has to be  accepted by the child type even if it doesn't need it.

	-Go allow functionality to be reused using a composition relationship

-Type Embedding
	-Type Embedding is they strategy that available to us provide a composition system in Go Language
*/

/*
-Polymorphism
	-with Polymorphism the consumer will not have information about  underlying type that provide functionality
	-the implementation details will be completely abstract from the consumer
	-ex:
		- Reader interface from Go Standard define one method Reader([]byte)(int,error)
		- we can read from  File , Tcp , WebSocket
		-So when writing an application that Consume Reader, we can pass any of these types(File,TCp,Socket)

-there is only one way to achieve Polymorphism in Go System,and it is  Interface

*/
import (
	"awesomeProject/OOP/Account"
	"awesomeProject/OOP/Payment"
	"awesomeProject/OOP/PaymentChannel"
	"fmt"
)

func main() {

	credit := Payment.CreateCreditAccount(
		"Mahmoud Reda",
		"1111-2222-3333-4444",
		5,
		2026,
		123)

	fmt.Printf("Owner name: %v\n", credit.OwnerName())
	fmt.Printf("Card number: %v\n", credit.CardNumber())
	fmt.Println("Trying to change card number")
	err := credit.SetCardNumber("invalid")
	if err != nil {
		fmt.Printf("That didn't work: %v\n", err)
	}
	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())

	/*Interface EnEncapsulation and abstract communication between sender and receiver using interface*/
	var option Payment.PaymentOption

	option = Payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	option.ProcessPayment(500)

	option = Payment.CreateCashAccount()

	option.ProcessPayment(500)

	//channel
	chargeCh := make(chan float32)
	PaymentChannel.CreateCreditAccount(chargeCh)
	chargeCh <- 500

	//composition
	ca := &Account.CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)

	//Composition solve delegation problem
	ha := &Account.HybridAccount{}
	fmt.Println(ha.AvailableFunds())
	fmt.Println(ha.CheckingAccount.AvailableFunds())
	var a string
	fmt.Scanln(&a)
}
