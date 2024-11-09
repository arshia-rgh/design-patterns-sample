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
