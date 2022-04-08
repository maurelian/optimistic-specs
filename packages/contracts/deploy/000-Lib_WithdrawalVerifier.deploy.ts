/* Imports: Internal */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import 'hardhat-deploy'

const deployFn: DeployFunction = async (hre) => {
  const { deploy } = hre.deployments
  const { deployer } = await hre.getNamedAccounts()

  await deploy('WithdrawalVerifier', {
    from: deployer,
    args: [],
    log: true,
    waitConfirmations: 1,
  })
}

deployFn.tags = ['Lib_WithdrawalVerifier']

export default deployFn
