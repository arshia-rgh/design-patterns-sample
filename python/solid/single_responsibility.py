class Journal:
    entries = []
    count = 0

    def add_entry(self, text):
        self.entries.append(f'{self.count}: {text}')
        self.count += 1

    def remove_entry(self, index):
        self.entries.pop(index)
        self.count -= 1

    # the persistence methods break the single responsibility
    def save_to_file(self, file_name, line_separator):
        pass

    def load_from_file(self, file_name):
        pass


# should separate the persistence (another class or just functions)
class Persistence:
    def __init__(self, line_separator):
        self.line_separator = line_separator

    def save_to_file(self, file_name):
        pass

    def load_from_file(self, file_name):
        pass
