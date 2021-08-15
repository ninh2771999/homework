function palindrome(str) {
    var re = /[^A-Za-z0-9]/g;
    str = str.toLowerCase().replace(re, '');
    var len = str.length;
    for (var i = 0; i < len/2; i++) {
      if (str[i] !== str[len - 1 - i]) {
          return false;
      }
    }
    return true;
   }
   console.log(palindrome("A man, a plan, a canal. Panama"));
   console.log(palindrome("eye"));
   console.log(palindrome("not a palindrome"));
   console.log(palindrome("0_0 (: /-\ :) 0-0"));