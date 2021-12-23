'use strict';
const Generator = require('yeoman-generator');
const chalk = require('chalk');
const yosay = require('yosay');

module.exports = class extends Generator {
  prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(
        `Welcome to the slick ${chalk.red('generator-gg-lambda-go')} generator!`
      )
    );

    const prompts = [
      {
        type: 'input',
        name: 'lambdaNameWithDash',
        message: 'Nombre de la lambda? En lugar de espacios use guiones',
        default: 'hello-world'
      }
    ];

    return this.prompt(prompts).then(props => {
      // To access props later use this.props.someAnswer;
      this.props = props;
    });
  }

  writing() {
    this._copyTemplate(this.props.lambdaNameWithDash)
    this._copySourceFiles(this.props.lambdaNameWithDash)
    this._copyOtherFiles()
    this.log("Copiado exitoso")
  }

  _copyTemplate(lambdaNameWithDash) {
    this.fs.copyTpl(
      this.templatePath('template.yaml'),
      this.destinationPath('template.yaml'),{
        lambdaNameWithDash: lambdaNameWithDash,
        lambdaNameCamelCase: this._toCamelCase(lambdaNameWithDash)
      }
    );
  }

  _toCamelCase(input){
    return input.replace('-','')
  }

  _copySourceFiles(lambdaNameWithDash){
   const files = ['main.go', 'main_test.go']
   files.forEach(f => this.fs.copy(
    this.templatePath(`lambda-name-with-dash/${f}`),
    this.destinationPath(`${lambdaNameWithDash}/${f}`)
  )  
  )
  }

  _copyOtherFiles(){
    const files = ['go.mod', 'go.sum','Makefile']
    files.forEach(f => this.fs.copy(
     this.templatePath(`${f}`),
     this.destinationPath(`${f}`)
   )  
   )
  }


};
