from abc import ABC, abstractmethod


class Creator(ABC):
    @abstractmethod
    def factory_method(self):
        # Can have extra logic here
        pass

    def some_operations(self):
        product = self.factory_method()

        return f"creator worked with {product.operation()}"


class Creator1(Creator):
    def factory_method(self):
        return Product1()


class Creator2(Creator):
    def factory_method(self):
        return Product2()


class Product(ABC):
    @abstractmethod
    def operation(self):
        pass


class Product1(Product):
    def operation(self):
        return "Result of Product1"


class Product2(Product):
    def operation(self):
        return "Result of Product2"


def client(creator: Creator):
    print(creator.some_operations())


if __name__ == "__main__":
    client(Creator1())
    client(Creator2())
