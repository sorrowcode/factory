# Factory

## Problem

- many structs share methods and maybe some attributes that are similar
- those structs have a similar purpose

```mermaid
classDiagram
	class BMW {
		+Drive()
		+OpenDoor(number)
	}
	class Mercedes {
		+Drive()
		+OpenDoor(number)
	}
	class VW {
		+Drive()
		+OpenDoor(number)
	}
```
- explicit instantiation is only with the direct struct possible

## Solution -> Factory Pattern

### Simple Example

- make all structs implementing an interface
- create a factory function which "produces" an object
	- sometimes this function needs an argument
	- sometimes the implementation itself is solving the differentiation

```mermaid
classDiagram
	class Car {
		<<interface>>
		+Drive()
		+OpenDoor(number)
	}
	class CarType {
		<<enumeration>>
		BMW 0
		VW 1
		Mercedes 2
	}
	class CarFactory {
		+createCar(carType) Car
	}
	class BMW {
		+BmwMethods()
	}
	class Mercedes {
		+MercedesMethods()
	}
	class VW {
		+VwMethods()
	}
	BMW ..|> Car
	Mercedes ..|> Car
	VW ..|> Car
	Car <.. CarFactory
	CarType -- CarFactory
```
- the created `Car` in `createCar(carType)` can than be casted dependend on the `CarType` if needed

## Explanation

- `BMW`, `Mercedes`, `VW` are called *concrete products*
- **clear code** want abstract dependencies
- therefore they implement from `Car` which is called *product interface*
- factory pattern would be nothing without a *factory*: `CarFactory`

```mermaid
classDiagram
	class ProductInterface {
		<<interface>>
		+someMethodToImplement()
	}
	class Factory {
		+createConcreteProduct() ProductInterface
	}
	ProductInterface <|.. ConcreteProductN
	ProductInterface <|.. ConcreteProduct1
	ProductInterface <.. Factory
```

## Implementation

Checkout the repository.

- modify in the `main.go` the different `CarType`s (0-2) to play with the functionality
	- 0 -> BMW
	- 1 -> Mercedes
	- 2 -> VW

