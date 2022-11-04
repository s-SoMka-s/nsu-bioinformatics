import sys

target_line = 6
val_pos = 4

def parse(file_name):
    lines = f.readlines()
    splitted = lines[target_line].split(" ")
    val = splitted[val_pos]
    result = val[1:]
    
    if float(result) > 90.0:
        print("OK")
    else:
        print("Not OK") 

if __name__ == '__main__':
    name = sys.argv[1]
    f = open(file_name, 'r')
    parse(f)
  