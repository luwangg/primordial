function [ minFactor ] = FindMinPrimeFactor(x)

  a = vpa((pi / 2 - 1 / x) * pi * x / (2 * 0.0068178951732107591857632195678137149876321876354431973), 100000)

  minFactor = vpa(floor(2 * a^.25), 100000);

  for i = -10000:1:10000

    if minFactor+i <= 0
      continue
    end

    if mod(x, minFactor+i) == 0
      minFactor = minFactor + i
      return
    end

  end

end

function [ factor ] = FindAFactor(a)

  primal = a^4 * 0.006817^2 / 2;

  for i = 1:50000000

    if mod(a, primal) == 0
        factor = primal;
        return
    end

    root = floor(sqrt(primal));

    if mod(a, root) == 0
        factor = root;
        return
    end

    primal = primal - root^2;

    if primal < 2
        factor = 1;
        return
    end

  end

end

function [ shape ] = FindShape2(x)

  shape = pi / 2 - 1 / (2 * x *x)

end

function [ shape ] = FindShape(x)

  digits = 2^29;

  shape = vpa(acos(vpa(sym(.5) / vpa(sym(x)*sym(x), digits) ) ), digits);

end

function [ sigma ] = FindSigma(a,b)

  digits = 2^29;

  shape = vpa(2 * acos(vpa(min(FindShape(a), FindShape(b)), digits) / vpa(max(FindShape(a), FindShape(b)), digits)), digits);

end

function [ largerPrime ] = FindLargerPrime( a, b )

  digits = 2^29;

  largerPrime = vpa(vpa(2 * FindSigma(a,b), digits) / vpa(abs(FindShape(b) - FindShape(a)), digits), digits)

end

function [ delta ] = FindDelta( a , b )

  digits = 2^29;

  delta = vpa(a * b / FindSigma(a,b), digits);

end

function [ primordial ] = FindPrimordial( a, b )

  digits = 2^29;

  primordial = vpa(0.006817 / abs(FindDelta(a,b)), digits);

end
