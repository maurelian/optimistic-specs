/* Imports: Internal */
import { DeployFunction } from 'hardhat-deploy/dist/types'
import 'hardhat-deploy'

const deployFn: DeployFunction = async (hre) => {
  const { deploy } = hre.deployments
  const { deployer } = await hre.getNamedAccounts()

  await deploy('L2OutputOracle', {
    from: deployer,
    args: [
      2,
      2,
      `0x${'00'.repeat(32)}`,
      0,
      '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
    ],
    log: true,
    waitConfirmations: 1,
  })
}

deployFn.tags = ['L2OutputOracle']

export default deployFn
