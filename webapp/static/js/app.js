// app.js

var app = angular.module('index', []);

app.controller('OrderCtrl', ['$scope','$http',
        function($scope, $http) {
            $scope.working = false;

            var logError = function(data, status) {
                console.log('Code '+status+': '+data)
                $scope.working = false;
            };
                

            $scope.search = function() {
            $scope.working = true;
            $http.get('/iElsa/workorder/'+$scope.workorder).
                error(logError).
                success(function(data) { 
                    $scope.orderlist = data.Workorder; 
                    $scope.working = false;
                });
            };
}]);
