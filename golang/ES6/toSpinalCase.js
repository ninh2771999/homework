function toSpinalCase(str) {

    // any string followed by upperCase letter
    var  re = /(?=[A-Z])/g;  
  
    // any string followed by space and upperCase/lowerCase letter 
    var  re2=/(?=\s[A-Z]|\s[a-z])/g;
  
    // any string of 2 or more '-'
    var re3 = new RegExp("-{2,}", "g");
  
    var space = new RegExp(" ","g");
    var hyphen = new RegExp("_","g");
  
    str = str.replace(hyphen,"-");
    str = str.replace(re, '-');   
    str = str.replace(re2,"-");
    str = str.replace(re3,"-"); 
    str = str.replace(space,"");
    str = str.toLowerCase();
    if(str.slice(0,1)== '-'){
      str = str.replace("-","");  
    }     
    return str;
  }
  
  let result = toSpinalCase('MyNameIsNinh');
  let result1 =toSpinalCase('my_name_is_Ninh_Handsome_Boi_Vip_Pro');
  console.log(result);
    console.log(result1);