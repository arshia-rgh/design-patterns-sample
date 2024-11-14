class Rectangle:
    def __init__(self, width, height):
        self._width = width
        self._height = height

    def get_width(self):
        return self._width

    def set_width(self, width):
        self._width = width

    def get_height(self):
        return self._height

    def set_height(self, height):
        self._height = height


"""
    Here the Square class inheritance from Rectangle but should override the setter methods to kkep it square
    and this way we will get unexpected result
    this breaks the liskov rule
"""


class Square(Rectangle):
    def __init__(self, size):
        super().__init__(size, size)

    def set_width(self, width):
        self._width = width
        self._height = width

    def set_height(self, height):
        self._width = height
        self._height = height


def use_it(rectangle: Rectangle):
    width = rectangle.get_width()
    rectangle.set_height(10)

    expected_area = 10 * width
    actual_area = rectangle.get_width() * rectangle.get_height()

    print(f"Expected: {expected_area}, actual: {actual_area}")


# Got unexpected for the square
rect = Rectangle(2, 3)
use_it(rect)  # Expected area: 20, Actual area: 20

sq = Square(5)
use_it(sq)  # Expected area: 50, Actual area: 100

# we should separate the square from the Rectangle for a solution to this
