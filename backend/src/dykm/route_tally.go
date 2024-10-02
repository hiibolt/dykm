package main;

func TallyResults ( req APIRequest) Result[APIResponse] {
    var result Result[APIResponse];
    
    switch (req.PIIType) {
      case "email":
        result.value.email += 0; //SnusbaseQuery(req.PII); Call to SnusbaseQuery

      case "phone":
        result.value.phone += 0; // BulkVSQuery(req.PII);

      case "username":
        result.value.username += 0;// SherlockQuery(req.PII);
        result.value.username += 0;// SnusbaseQuery(req.PII);

      case "name":
        result.value.name += 0; // SnusbaseQuery(req.PII);

      case "ip":
        result.value.ip += 0; // SnusbaseQuery(req.PII);
        result.value.ip += 0; // SnusbaseGeo(req.PII);

      case "hash":
        result.value.hash += 0; // SnusbaseQuery(req.PII);
        result.value.hash += 0; // SnusbaseHashing(req.PII);

      case "password":
        result.value.password += 0; // SnusbaseQuery(req.PII);

      default:
        break;
    } 

    return result;
}