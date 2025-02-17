# Formulas

This directory contains the formulas used to calculate the distances between points on a sphere.

## Haversine

The haaversine formula is based on the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula) and is used to calculate the distance between two points on a sphere. The formula is:

$$\\Delta \sigma = 2\cdot arcsin\sqrt{sin^2\left (\frac{\Delta\phi}{2}\right )+\left (cos^2\left(\frac{\phi_1 + \phi_2}{2}\right ) - sin^2\left (\frac{\Delta\phi}{2}\right ) \right) \cdot sin^2\left (\frac{\Delta\lambda}{2}\right )}\$$

## Vincenty

The Vincenty formula is based on the [Great-circle distance](https://en.wikipedia.org/wiki/Great-circle_distance) and is used to calculate the distance between two points on a sphere. Technically, it is not the Vincenty formula, but it is what's used for my CS314 class. The formula is:

$$\\Delta \sigma = arctan \frac{\sqrt{\left (cos\phi_2 \cdot sin\left (\Delta\lambda\right ) \right)^2 + \left (cos\phi_1 \cdot sin\phi_2 - sin\phi_1 \cdot cos\phi_2 \cdot cos\left (\Delta\lambda\right ) \right)^2}}{sin\phi_1 \cdot sin\phi_2 + cos\phi_1 \cdot cos\phi_2 \cdot cos \left (\Delta\lambda\right )}\$$

## Spherical Law of Cosines

The Spherical Law of Cosines formula is based on the [Great-circle distance](https://en.wikipedia.org/wiki/Great-circle_distance) and is used to calculate the distance between two points on a sphere. The formula is:

$$\\Delta \sigma = arccos\left(sin\phi_1 \cdot sin\phi_2 + cos\phi_1 \cdot cos\phi_2\ \cdot cos\Delta\lambda\right)\$$

## Lastly...

Once $\Delta \sigma$ is calculated, the distance between the two points is simply $d = r \Delta\sigma\$, where $r$ is the radius of the sphere and $\Delta \sigma$ is the calculated distance.

### Note:

If any of the calculations are incorrect, please let me know and I'll fix them.
