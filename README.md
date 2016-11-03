# Primordials

A primordial is a certain type of number that contains all the primes less than one half the number as approximate factors of it.  Examples include 69, 1574, and 220085.  The equations to calculate the primordial stem from another type of calculation I discovered that is used to calculate another prime from component
primes, which always recovers a prime number (or is one integer value off due to floating-point errors). So, the following equations are true:

# The Shape of A Number

The shape of a number is a new property of a number that I've discovered that I'm still analyzing and coalescing but here is the equation I used to calculate the shape:

      1/2 = x*x*cos( shapeX )

I think the appropriate letter to use for this is the lowercase phi φ so:

      φ(x) = acos( .5 / x^2 )

Additional Work (2 Nov 2016):

I have done more work on shapes of non-prime numbers and have found the following formula to be a more general form of the shape of a number:

      1/2^(n*u) = ( Γ b(i) )^2 cos φ(x)

Where n is the total number of prime factors in the number x, u is the number of unique factors in the number x, Γ b(i) is the product of all the basic factors of a prime number, meaning the unique prime factors and the number 1.  So, for example, the shape of the number 63 is as follows:

      1/2^(3*2) = ( 1 * 3 * 7 )^2 * cos φ(63)

# Some Interesting Properties of 1 and 0

So, some manipulation of the shape of a number formula yields some interesting properties.  One, is that the number 1 does not work in the formula, as it is neither prime, nor composite, and has only a single basic factor, itself.  Plugging  it into the equation yields a shape of zero, which may or may not be accurate but I've found the shape of one to be approximately 0.006817 but that number may have a different name than shape, so further research is necessary.

Using some manipulation, we can see there are negative values possible in the shape of a number formula, so negative shapes are feasable.  Meaning we should be able to factor negative numbers into their basic primes as well, opening up the entire range of imaginary and real numbers to shapes and factorization and prime number theory.

## Shape of Zero

I assumed the cosine of some shape was equal to zero, and created the following formula:

      0 = 1 / (2 ^ (n*u) * (Γ b(i) )^2 )

And then solved for values, indicating the following: 0 has an infinite number of prime factors, both unique and redundant, but does not have itself as a basic factor.  This is required for the equation to evaluate to true.

This suggests the following: All numbers are factors of zero, for it to have infinitely many factors that are both unique and not unique.  This suggests that ALL numbers are factors of zero, meaning, if you multiply EVERY number together, except zero, you get zero.  This implies that zero is not a number as well, as it is not a basic factor, nor a prime factor, or the equation would evaluate to a division by zero.  Either zero is not a number, or dividing by zero yields positive and negative infinity, combined. (zinfinity?)

So, either way, interesting!  I feel that it is impossible to tell without further research whether zero is a basic factor of itself or not, but it does have a shape, suggesting it is a number.

## Shape of One and Negative One

We can also solve for cosine equal to 1 and -1, and get the following:

      1 = 1 / (2 ^ (n*u) * (Γ b(i) )^2 )

And we see the only way for this to evalue to true is for the denominator to be equal to 1, so we solve for that and find that 1 has zero prime factors and zero unique prime factors, but does contain itself as a basic factor.  Therefore, one is a factor of itself, but is neither prime nor composite.

## Factoring Negative Numbers

It also shows that 1 is a factor of every number, both positive and negative, and negative one is a factor of every negative number, using a similar method.  This allows us to factor all negative numbers into composing basic factors.

It also shows that you cannot use the negative basic factors of a positive number's composing basic factors to 'simulate' the factors of a composite number.  I.e. you cannot say -1, 1 and -2 are factors of +2 because there is only 1 prime factor and 1 unique prime factor of +2.

There is, however, an interesting thought I've just had, that if you're counting the number of negative factors of a positive number you merely flip the signs on both the number of prime factors and the number of unique prime factors, for example, n = -1 and u = -1, then the equation works out the same, with -1 and -2 being the basic factors (and -2 being the prime factor) but this line of thinking I think will prove false, as I'm using the actual count of prime and unique prime factors in my equation, so, just a thought but I believe my earlier proof still stands.

## Aeon's Number

It is with this knowledge, that the actual shape of 1 is zero, and the shape of zero is +-pi/2, that I move to rename my "ShapeOfOne" variable with 0.006817 as its approximate value to Aeon's Number, using the uppercase letter A to signify it.

# Several New Trigonometric Approximations

## cosine and inverse cosine

Stemming from this recent work, I've also discovered a new approximation to sin, cos, tan, arcsin, arccos, and arctan which I'll describe here:

I began with attempts to find sigma of a composite number as follows:

      ( ( 2^(n*u) * ( Γ b(i) )^2 ) ^ 2 - 1 ) ^ 0.5 / ( 2^(n*u) * ( Γ b(i) ) ^ 2 )

And plugged it into the shape of a number formula to get the following:

      acos( ... ) * 2^(n*u) * ( Γ b(i) )^2 ~= 1

Where the ... is the formula directly above.

From this equation, I derived the following approximations for cos and acos:

      acos( sqrt( a ^ 2 - 1) / a ) ~= 1 / a
      sqrt( a ^ 2 - 1 ) / a ~= cos( 1 / a )

## sine and inverse sine

And the following for sin and asin:

      asin( 1/a ) ~= sin( 1/a ) ~= 1/a

The limit on a is that it must be greater than 1 and the approximations approach the actual values as a approaches infinity.

## tangent and inverse tangent

Here is another trigonometric approximation that I found while trying to solve my equations, using a bit of geometry and the Pythagorean theorem:

      atan( 1 / b - 1 / a ) ~= tan( 1 / b - 1 / a ) ~= 1 / b - 1 / a

## possible uses in computation and calculus

These new trignometric approximations allow for very, very rapid approximation of the trigonometric functions, with exact accuracy possible using larger values, allowing for a very finely graduated curve of accuracy vs computation time.

Another possible use is in the calculation of derivatives and integrals with these formulae in them by allowing for substitution and approximation.  You would not even need to use a u substitution, except for solving, perhaps, but the trigonometric function would be a direct one-to-one plug in (or vice versa).

# The Sigma of A Number

The sigma of a number is related to the shape of two prime numbers combined, and has the following equation for its calculation:

      ς(a,b) = 2 * acos( min(φ(a), φ(b)) / max(φ(a), φ(b)) )

Additional Work (2 November 2016):

Some additional work on the sigma of a composite number has yielded me the following equation:

      ς(x) = ( 2 * x ) / ( A * ( Γ b(i) ) ^ 2 * 2 ^ (n * u) )

This is the current area of research I'm working on right now, however.

# Finding Additional Prime Numbers

To find additional prime numbers, we're going to manipulate the sigma value a bit as follows:

      AnotherPrime = round( 2 * ς(a,b) / abs( φ(b) - φ(a) ) )

There is a strange pattern to this function, as it will always return a decimal number, needing to be rounded to an integer value to be used as a prime, and then it has a strange method of returning a 5 in the 1's digit if both the 3 and 7 are prime, and sometime returning an even number (which should be either ceiled or floored as appropriate).  However, I've yet to discover a single result that when adjusted slightly as indicated does not yield a prime number, though I have not yet found
a pattern to the calculated primes - sometimes the prime is between a and b and sometimes it is greater than a or b.

Additionally, I've not tried combining the shapes of non-primes or imaginary numbers or a slew of other possible values that could yield further interesting results.

# Finding the Delta of A Primordial

This uses the sigma of two numbers in a slightly different way, and is the foundation for finding a primordial for any two prime components:

      Δ(a,b) = (a * b) / ς(a,b)

# Finding a Primordial

To find the actual primordial, we manipulate the delta value as follows:

      Primordial = φ(1) * Δ(a,b)

      CORRECTION:

      Primordial = A * Δ(a,b)

Where φ(1) (or A) is explained below.

# Shape of One - Renamed to Aeon's Number "A"

In order to calculate the shape of a number, I originally started by finding the angle between two numbers, using the modified Pythagorean formula.  I then found that in order to find the shape of a standalone number, you needed to take the shape between zero and the number, however, to take the shape between zero and the number is impossible as the equation would collapse into a simple 0 = 0 or similar.

Therefore, I found the angle between 1 and any number, and then I used experimentation to find the common factor between them, which comes out to the following approximation:

        φ(1) = 0.006817

        CORRECTION: I have found the true shape of 1 to be zero, and the shape of zero to be +- pi/2 so I'm renaming this new constant to be Aeon's Number, or an uppercase A.

        A ~= 0.006817

  A different equation was used to calculate this value, and can be calculated to many more digits, however, I've found this approximation is very useful.  See the section below for more details.

# Miscellaneous Equations

Here I'm including some extra equations I used and intermediate formulae that I'm working on for such things as decomposing composite numbers into primes and work like that.  These equations are still works in progress where indicated and I'm hoping other people will be able to further this research along!  Thanks in advance!

    * Shape of a Composite Number (used to calculate ShapeOfOne):

      (ab) ^ 2 = a^2 + b^2 - 2 * (ab)^2 * cos(shapeAB)

      where shapeAB is the angle between the numbers a and b

    * Another form of the Shape of a Number formula

      1 = a - 2*cos(shapeA)

    * Work On Factoring Component Numbers Into Prime Factors

      x = acos( 1 / (2b^2) ) - acos( 1 / (2a^2) ) + ( acos( 1 / (2a^2 ) / acos( 1 / (2b^2) ) ) )
      x = atan( ( b - a ) / 2 ) + cos((a^2) / (b^2))

    * Some Alternate Equations I'm Working On

      Theta(p) = 0.5 * a * b - 2 * p^2
      p(a,b) = 0.5 * a + 0.5 * b + 2 * (ab)^2
      Theta(a) = p(a,b) + 2*a*b*cos(Theta(a))
      Theta(b) = p(a,b) + 2*a*b*cos(Theta(b))

      Theta(a,b) = 2*Theta(a) + 2*Theta(b) - atan((ab)^2)

    * Finally, Even More Equations I'm Working On As Of This Very Moment

      Theta(a,b) = 4a + 4b + 2ab * cos(thetaA) * cos(thetaB)

      Warning! I know these are wrong, and have a modified version beneath them.
      Theta(a) = 5*a - sin(thetaA) + a * ShapeOfOne
      Theta(b) = 5*b - sin(thetaB) + b * ShapeOfOne

      ThetaA - 3.5 * a * ShapeOfOne * sin(ThetaA) = X / ShapeOfOne

      X is a number I've found that is related somehow to the value A, I saw it
      on my calculator and remembered it but had already cleared the screen OOPS!

      ThetaAB = 5*a + 5*b - 3absin(thetaA * thetaB)

    * Here is an additional trigonometric identity that is a great approximation, but spawned from my work so I'd like it stored somewhere in case someone can make sense of it:

        1 / b - 1 / a = tan( 1 / b - 1 / a )

# Perfect Encryption Using Primordials (a conjecture)

Primordials are used in the perfect encryption method by using a technique as follows: two or more primordials can be combined to contain multiple sets of primes from various ranges.  Therefore, it is impossible to know exactly which primordials were used, as all prime factors are combined when the primordials are multiplied together, or perhaps combined in some other manner, and since the prime factors are approximate, if you were to select them knowingly (i.e. you had the key), you could determine the two or more primordials used to compose the combined integer, but without the key, you'd be left with a statistics problem trying to figure out which primes were used (because all prime factors are varying levels of approximate, instead of integer values) and as a result, the encryption should be unbreakable without the key.

Since we can combine infinitely many primordials with infinitely many digits, limited only by the space you wish to store the key in, and we can generate them instantly, this encryption technique should be perfect.

Just a thought, however, as I've not the time right now to write an implementation, though the implementation should stem pretty readily from an RSA technique using the known prime factors of the primordials you choose.

So the 'secret' so to speak is exactly which primordials you used, or perhaps, which prime numbers you used to generate the primordials in the first place, as, as far as I know, there is no way to reverse the primordial and discover the prime factors, as it has many prime factors, and multiple pairs of primes generate the same primordial.  So basically, you can find new primes to use by recursing several simple primes many many times, using my "Find Another Prime" function above, then use those primes to find one or more primordials, and perhaps use the recurrance of the prime factors as the values to encrypt/decrypt the message using public key/private key or the primes themselves and incorporate the primordials in some other way.

# Some Example Work

I recently ported the code to Matlab (you'll see the scripts above), however, I hit the variable-precision arithmetic limit quite quickly.  I need a much higher precision arc cosine function, so perhaps I'll work on that next.  But here are a few of the prime numbers I generated before running out of precision:

1,751,689,182,900,215,667

13,478,474,172,247,698,097

11,276,769,660,352,321,127

# Copyright, Improvements, and Using For Your Own Devices

This is as yet unpublished work, however, I am publishing it now on the internet through this Github account in order to seek comment, enlightenment, and opinion on my work, so feel free to borrow it, improve it, and use it in your work, though proper credit to me is required.  And I would appreciate if you resubmit any improvements to me so that I may continue to expand this compendium on this new field of mathematics.

Copyright 2016 Chris Pergrossi

And this technique was discovered by myself only on 26 October 2016 and 27 October 2017, with additional help given by Alex Mcghee during the proving stage.

Thanks!
