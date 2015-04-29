// app.js

var app = angular.module('index', []);

app.controller('OrderCtrl', ['$scope','$http',
        function($scope, $http) {
            $scope.working = false;
            $scope.showed = false;
            $scope.notfound = false;

            var logError = function(data, status) {
                console.log('Code '+status+': '+data)
                $scope.working = false;
                $scope.notfound = true;
            };


            $scope.search = function() {
                $scope.showed = false;
                $scope.working = true;
                $scope.notfound = false;
                if ($scope.searchtype == 'hd num') {
                    $http.get('/iElsa/api/hdnum/'+$scope.dig).
                        error(logError).
                        success(function(data) { 
                            $scope.showed = true;
                            console.log(data);
                            $scope.orderlist = data.Workorder; 
                            $scope.working = false;
                        });
                }

                if ($scope.searchtype == 'zip name') {
                    $http.get('/iElsa/api/file/'+$scope.dig).
                        error(logError).
                        success(function(data) { 
                            $scope.showed = true;
                            console.log(data);
                            $scope.orderlist = data.Workorder; 
                            $scope.working = false;
                        });
                }

                if ($scope.searchtype == 'work order') {
                    $http.get('/iElsa/api/order/'+$scope.dig).
                        error(logError).
                        success(function(data) { 
                            $scope.showed = true;
                            console.log(data);
                            $scope.orderlist = data.Workorder; 
                            $scope.working = false;
                        });
                }

            };
        }]);
