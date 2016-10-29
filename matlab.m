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
