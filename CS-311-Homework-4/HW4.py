VERYSMALL = 0.00000001
def findSlope(a, b, c, x) :
    return 2*a*x + b

def findZero(a, b, c, x) :
    fn_x = a*x*x + b * x + c
    if abs(fn_x) < VERYSMALL:
        return x
    slope_x = findSlope(a, b, c, x)
    next_guess = x - fn_x/slope_x
    return findZero(a, b, c, next_guess)

# calling code

findZero(1, -2, 1, 5);