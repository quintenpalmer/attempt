#!/bin/python

import os

def main():
    packages = os.listdir(os.path.join(os.getcwd(), "src"))
    packages.remove("github.com")
    packages.remove("cgl.tideland.biz")
    os.system("go test " + " ".join(packages))

if __name__ == '__main__':
    main()