var galleryApp = angular.module('galleryApp', []);

galleryApp.controller('galleryController', function galleryController($scope, $http) {
    $scope.imageList = null;

    //Get image list
    $http.get("./images").then(function(response) {
        $scope.imageList = response.data.images;
    });
});