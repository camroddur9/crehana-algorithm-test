from sys import stdin

def main():
    billsCount = [int(x) for x in stdin.readline().split(",")]
    bills25 = 0
    bills50 = 0
    bills100 = 0
    for i in range(len(billsCount)):
        if(billsCount[i] == 25):
            bills25 += 1
        elif(billsCount[i] == 50):
            if(bills25 > 0):
                bills25 -= 1
                bills50 += 1
            else: 
                return("NO")
        elif(billsCount[i] == 100):
            if(bills25 > 0 and bills50 > 0):
                bills25 -= 1
                bills50 -= 1
                bills100 += 1
            elif(bills25 > 2):
                bills25 -= 3
                bills100 += 1
            else:
                return("NO")
    return("YES")

print(main())