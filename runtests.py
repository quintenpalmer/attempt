#!/bin/python

import os

def main():
    packages = os.listdir(os.path.join(os.getcwd(), "src"))
    os.system("go test " + " ".join(packages))

if __name__ == '__main__':
    main()