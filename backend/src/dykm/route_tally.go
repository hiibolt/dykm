package main;

func TallyResults ( req APIRequest) Result[Tally] {
    var result Tally;
    
    switch (req.PIIType) {
      case "email":
        result.Emails += 0; //SnusbaseQuery(req.PII); Call to SnusbaseQuery

      case "phone":
        result.Phones += 0; // BulkVSQuery(req.PII);

      case "username":
        result.Usernames += 0;// SherlockQuery(req.PII);
        result.Usernames += 0;// SnusbaseQuery(req.PII);

      case "name":
        result.Names += 0; // SnusbaseQuery(req.PII);

      case "ip":
        result.Ips += 0; // SnusbaseQuery(req.PII);
        result.Ips += 0; // SnusbaseGeo(req.PII);

      case "hash":
        result.Hashes += 0; // SnusbaseQuery(req.PII);
        result.Hashes += 0; // SnusbaseHashing(req.PII);

      case "password":
        result.Passwords += 0; // SnusbaseQuery(req.PII);
        result.Passwords += 0; //SnusBaseHashing(req.PII);

      default:
        break;
    } 

    return Ok(result);
}