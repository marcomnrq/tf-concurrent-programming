import http from './http-common';

class ApiService {
    getTemperature(data){
        return http.post('/getPrediction', data)
    }
}

export default new ApiService();