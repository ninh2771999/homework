let arr1 = [3, 5, 2, 2, 5, 5];
let arr2 = [2, 1, 66, 5];
let arr3 = [4, 1, 66, 8];
let unique = [...new Set([...arr1,...arr2,...arr3])];
console.log(unique);
// function arrayUnique(array) {-----------------> second way
//     var a = array.concat();
//     for(var i=0; i<a.length; ++i) {
//         for(var j=i+1; j<a.length; ++j) {
//             if(a[i] === a[j])
//                 a.splice(j--, 1);
//         }
//     }

//     return a;
// }

// var array1 = [1,2,5,8];
// var array2 = [3,5,7,8];
//     // Merges both arrays and gets unique items
// var array3 = arrayUnique(array1.concat(array2));
// console.log(array3)