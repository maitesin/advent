import sys
from textwrap import wrap

class Layer:
    def __init__(self, data, width, height):
        self._data = data
        self._width = width
        self._height = height
    
    def number_of_zero_digits(self):
        return self._data.count('0')
    
    def checksum(self):
        return self._data.count('1') * self._data.count('2')

    def __len__(self):
        return len(self._data)
    
    def __repr__(self):
        return 'Layer(data={}, width={}, height={})'.format(self._data, self._width, self._height)

class Image:
    def __init__(self, layers):
        self._layers = layers
    
    def checksum(self):
        min_number_of_zeros = sys.maxsize
        min_layer = None
        for layer in self._layers:
            zeros = layer.number_of_zero_digits()
            if zeros < min_number_of_zeros:
                min_number_of_zeros = zeros
                min_layer = layer
        return min_layer.checksum()
                
    
    def __repr__(self):
        return 'Image(layers={})'.format(self._layers)


def main(filename, width, height):
    with open(filename, 'r') as f:
        chunk_size = width * height
        data = f.readline()
        chunks = wrap(data, chunk_size)
        layers = [Layer(chunk, width, height) for chunk in chunks]
        image = Image(layers)
        print(image.checksum())

if __name__ == "__main__":
    main(sys.argv[1], int(sys.argv[2]), int(sys.argv[3]))