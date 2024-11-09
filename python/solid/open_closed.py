from abc import abstractmethod, ABC


class Product:
    def __init__(self, color, size):
        self.color = color
        self.size = size


# In Filter class if we need to add another filters like (filter_by_color_and_size)
# we need to modify the Filter class and add new methods, and this will break the OCP
class Filter:
    def filter_by_color(self, color, products):
        result = []
        for i in products:
            if i.color == color:
                result.append(i)
        return result

    def filter_by_size(self, size, products):
        result = []
        for i in products:
            if i.size == size:
                result.append(i)
        return result


# Used specification pattern
# We use new class for each new fiter that should implement the same interface
class Specification(ABC):
    @abstractmethod
    def is_satisfied(self, product):
        pass


class ColorSpecification(Specification):
    def __init__(self, color):
        self.color = color

    def is_satisfied(self, product):
        return product.color == self.color


class SizeSpecification(Specification):
    def __init__(self, size):
        self.size = size

    def is_satisfied(self, product):
        return product.size == self.size


class AndSpecification(Specification):
    def __init__(self, *specs):
        self.specs = specs

    def is_satisfied(self, product):
        return all(spec.is_satisfied(product) for spec in self.specs)


# Now this Filter class doesn't break the OCP and use specification for filters
#  if a new filter is added, we need to implement a new class and not to modify the existing one

class BetterFilter:
    def filter(self, spec, products):
        return [product for product in products if spec.is_satisfied(product)]
