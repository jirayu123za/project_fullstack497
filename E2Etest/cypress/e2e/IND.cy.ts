
describe('Login Functionality Test', () => {
  beforeEach(() => {
    cy.visit('http://localhost:5173/landing')  
  });



  it('should successfully log in with valid credentials', () => {
    cy.get('input[id="username"]').type('username1');
    cy.get('input[id="password"]').type('password');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/insdash'); 
  });

});



describe('Course Creation Functionality Test', () => {
  beforeEach(() => {
    cy.visit('http://localhost:5173/landing'); 
    cy.get('input[id="username"]').type('username1');
    cy.get('input[id="password"]').type('password');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/insdash');
  });

  it('should successfully create a new course with cover image url', () => {
    cy.get('img[src="/src/icons/ion_create.png"]').eq(1).click(); 
    cy.url().should('include', '/create');
    cy.get('input[id="course_name"]').type('Full Stack Development');
    cy.get('input[id="course_code"]').type('CSD101'); 
    cy.get('input[id="img_url"]').type('https://images2.minutemediacdn.com/image/upload/c_crop,w_2560,h_1440,x_0,y_0/c_fill,w_1200,ar_16:9,f_auto,q_auto,g_auto/images/voltaxMediaLibrary/mmsport/video_games/01hxvjzw09p2wwtdz34w.jpg'); // ใส่ URL ของรูปภาพ
    cy.get('select[id="pickcolor"]').select('Purple'); 
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/insdash');
    cy.contains('Full Stack Development');
  });
});




describe('Assignment Creation Functionality Test', () => {
  beforeEach(() => {
    cy.visit('http://localhost:5173/landing'); 
    cy.get('input[id="username"]').type('username1');
    cy.get('input[id="password"]').type('password');
    cy.get('button[type="submit"]').click();
    cy.url().should('include', '/insdash');
  });

  it('should successfully create a new assignment in the first course', () => {
    cy.contains('test01').click();
    cy.url().should('include', '/course/');
    cy.get('div').contains('Assignment').parent().next().click();
    cy.contains('Add New Assignment').should('be.visible');
    cy.get('input:not([readonly])').eq(0).clear().type('Assignment 1');
    cy.get('input:not([readonly])').eq(1).clear().type('Description for Assignment 1'); 
    cy.get('input[type="date"]').clear().type('2024-10-30');
    cy.get('button').contains('Add Assignment').click();
    cy.contains('Assignment 1').should('exist');
    cy.reload();
    cy.contains('Assignment 1').should('exist');
  });
});








