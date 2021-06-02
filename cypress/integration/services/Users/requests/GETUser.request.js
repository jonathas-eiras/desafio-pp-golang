/// <reference types="cypress" />

const token = Cypress.env('token')

function getUser(idUser) {
  return cy.request({
    method: 'GET',
    url: `/users/${idUser}`,
    auth:{
          bearer: token
        } 
  })
}

export { getUser };