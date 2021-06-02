/// <reference types="cypress" />

const token = Cypress.env('token')

function deleteUSer(idUser) {
    return cy.request({
        method: 'DELETE',
        url: `/users/${idUser}`,
        auth:{
                bearer: token
            } 
      })
  }
  
  export { deleteUSer };
  