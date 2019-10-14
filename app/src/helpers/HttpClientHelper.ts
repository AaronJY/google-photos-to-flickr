export class HttpClientHelper {
    static objectToQuery(obj: {}): string {
        const objKeys = Object.keys(obj);
        if (!objKeys)
            return "";

        let query = "?";
        objKeys.forEach((val, i) => {
            query += `${val}=${obj[val]}`;
            if ((i + 1) < objKeys.length)
                query += "&";
        });

        return query;
    }
}