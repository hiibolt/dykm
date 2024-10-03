package main;

func TallyResults ( req APIRequest) Result[Tally] {
    var result Tally;
    
    switch (req.PIIType) {
      case "email":
        tally_result := SnusbaseQuery(req.PIIType, req.PII);
        
        if tally_result.IsErr(){
          return tally_result;
        }
        var tally Tally = tally_result.UnwrapOk();
        result.add(tally);

      case "phone":
        tally_result := BulkVSQuery(req.PII);

        if tally_result.IsErr(){
          return tally_result;
        }

        var tally Tally = tally_result.UnwrapOk();
        result.add(tally);

      case "username":
        tally_result1 := SnusbaseQuery(req.PIIType, req.PII);
        tally_result2 := SherlockQuery(req.PII);

        if tally_result1.IsErr(){
          return tally_result1;
        }

        if tally_result2.IsErr(){
          return tally_result2;
        }

        //Unwrap both of the tallies
        var tally1 Tally = tally_result1.UnwrapOk();
        var tally2 Tally = tally_result2.UnwrapOk();

        //Add it to the final result
        result.add(tally1);
        result.add(tally2);

      case "name":
        tally_result1 := SnusbaseQuery(req.PIIType, req.PII);

        if tally_result1.IsErr(){
          return tally_result1;
        }

        var tally1 Tally = tally_result1.UnwrapOk();
        result.add(tally1);

      case "ip":
        tally_result1 := SnusbaseQuery(req.PIIType, req.PII);
        tally_result2 := SnusbaseGeo(req.PII);

        if(tally_result1.IsErr()){
          return tally_result1;
        }

        if(tally_result2.IsErr()){
          return tally_result1;
        }
        
        //Unwrap both tallies
        var tally1 Tally = tally_result1.UnwrapOk();
        var tally2 Tally = tally_result2.UnwrapOk();
        
        //Add it to the final result
        result.add(tally1);
        result.add(tally2);

      case "hash":
        tally_result1 := SnusbaseQuery(req.PIIType, req.PII);
        tally_result2 := SnusbaseHashing(req.PIIType, req.PII);

        if(tally_result1.IsErr()){
          return tally_result1;
        }

        if(tally_result2.IsErr()){
          return tally_result1;
        }
        
        //Unwrap both tallies
        var tally1 Tally = tally_result1.UnwrapOk();
        var tally2 Tally = tally_result2.UnwrapOk();
        
        //Add it to the final result
        result.add(tally1);
        result.add(tally2);

      case "password":
        tally_result1 := SnusbaseQuery(req.PIIType, req.PII);
        tally_result2 := SnusbaseHashing(req.PIIType, req.PII);

        if(tally_result1.IsErr()){
          return tally_result1;
        }

        if(tally_result2.IsErr()){
          return tally_result1;
        }
        
        //Unwrap both tallies
        var tally1 Tally = tally_result1.UnwrapOk();
        var tally2 Tally = tally_result2.UnwrapOk();
        
        //Add it to the final result
        result.add(tally1);
        result.add(tally2);

      default:
        break;
    } 

    return Ok(result);
}