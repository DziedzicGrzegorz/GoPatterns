#include <string>
#include <cstdlib>

// Helper method to calculate the Greatest Common Divisor (GCD) using Euclid's algorithm
int gcd(int a, int b) {
    while (b != 0) {
        int t = b;
        b = a % b;
        a = t;
    }
    return a;
}

extern "C" {
    const char* fractionAddition(const char* expression) {
        static std::string result;  // Static to persist after function return

        // Initialize numerator and common denominator which is lcm of 6, 7, 8, 9, 10
        int numerator = 0;
        int commonDenominator = 6 * 7 * 8 * 9 * 10;

        std::string expr(expression);

        // Add '+' to the beginning if the expression starts with a number
        if (isdigit(expr[0])) {
            expr = "+" + expr;
        }

        size_t i = 0;
        size_t lengthOfExpression = expr.size();

        // Iterate through the characters in the expression
        while (i < lengthOfExpression) {
            // Determine the sign of the term (+1 or -1)
            int sign = expr[i] == '-' ? -1 : 1;
            i++;  // Move past the sign for parsing the number

            // Determine the position of the next sign (or end of string)
            size_t j = i;
            while (j < lengthOfExpression && expr[j] != '+' && expr[j] != '-') {
                j++;
            }

            // Extract the fraction and split it into numerator and denominator
            std::string fraction = expr.substr(i, j - i);
            size_t slashPosition = fraction.find('/');
            int currentNumerator = std::stoi(fraction.substr(0, slashPosition));
            int currentDenominator = std::stoi(fraction.substr(slashPosition + 1));

            // Adjust the numerator according to the common denominator and the sign
            numerator += sign * currentNumerator * commonDenominator / currentDenominator;

            // Move the index 'i' to the position of the next term's sign or the end of the string
            i = j;
        }

        // Simplify the fraction by dividing both numerator and denominator by their GCD
        int greatestCommonDivisor = gcd(abs(numerator), commonDenominator);
        numerator /= greatestCommonDivisor;
        commonDenominator /= greatestCommonDivisor;

        // Store the result in a static string
        result = std::to_string(numerator) + "/" + std::to_string(commonDenominator);
        return result.c_str();
    }
}
