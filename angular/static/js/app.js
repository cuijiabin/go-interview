// Declare app level module which depends on filters, and services
angular.module('angular-go', ['ngResource', 'ngRoute', 'ui.bootstrap', 'ui.date'])
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: '/static/views/home/home.html',
        controller: 'HomeController'})
      .otherwise({redirectTo: '/'});
  }]);
